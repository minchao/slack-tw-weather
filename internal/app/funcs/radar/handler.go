package radar

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"time"

	"github.com/minchao/slack-tw-weather/internal/pkg/slack"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/disintegration/imaging"
)

const (
	sourceImageURLFormat = "http://www.cwb.gov.tw/V7/observe/radar/Data/HD_Radar/CV1_3600_%s.png"
	imageURLFormat       = "https://s3-%s.amazonaws.com/%s/%s"
)

var (
	local *time.Location

	region = os.Getenv("AWS_REGION")
	bucket = os.Getenv("RADAR_BUCKET")
)

func init() {
	local, _ = time.LoadLocation("Asia/Taipei")
}

func Handler(_ context.Context, snsEvent events.SNSEvent) {
	record := snsEvent.Records[0]
	snsRecord := record.SNS
	fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)

	imageURL, dateTime, err := fetchRadarImageURL()
	if err != nil {
		fmt.Println("fetch image error:", err)
		return
	}
	t, _ := time.Parse("200601021504", dateTime)

	message := slack.Message{
		ResponseType: "in_channel",
		Text:         "雷達回波圖",
		Attachments: []slack.Attachment{
			{
				Text:     t.Format("2006/01/02 15:04"),
				ImageURL: imageURL,
			},
		},
	}
	messageBytes, _ := json.Marshal(message)
	_, err = http.Post(snsRecord.Message, "application/json", bytes.NewBuffer(messageBytes))
	if err != nil {
		fmt.Println("send error:", err)
	}
}

func generateImageNameSuffix(d time.Duration) string {
	now := time.Now().In(local)
	t := now.Add(d).Format("200601021504")
	t = t[:len(t)-1] + "0"
	return fmt.Sprintf("%s", t)
}

func fetchRadarImageURL() (imageURL string, dateTime string, err error) {
	svc := s3.New(session.Must(session.NewSession()))
	dateTime = generateImageNameSuffix(-time.Minute * 10)
	filename := fmt.Sprintf("%s.jpg", dateTime)

	_, err = svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err == nil {
		return fmt.Sprintf(imageURLFormat, region, bucket, filename), dateTime, nil
	}

	source, _, err := fetchSourceImage(dateTime)
	if err != nil {
		// Retry
		dateTime = generateImageNameSuffix(-time.Minute * 20)
		source, _, err = fetchSourceImage(dateTime)
		if err != nil {
			return "", "", errors.New("failed to fetch radar image")
		}
	}
	thumbnail, err := createThumbnail(source)
	if err != nil {
		return "", "", errors.New("failed to create thumbnail")
	}

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		ACL:    aws.String("public-read"),
		Body:   bytes.NewReader(thumbnail.Bytes()),
		Metadata: map[string]*string{
			"Key": aws.String("MetadataValue"),
		},
	})
	if err != nil {
		return "", "", err
	}

	return fmt.Sprintf(imageURLFormat, region, bucket, filename), dateTime, nil
}

func fetchSourceImage(dateTime string) (image.Image, string, error) {
	resp, err := http.Get(fmt.Sprintf(sourceImageURLFormat, dateTime))
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	return image.Decode(resp.Body)
}

func createThumbnail(img image.Image) (*bytes.Buffer, error) {
	var imgBuffer bytes.Buffer
	thumbnail := imaging.Thumbnail(imaging.CropCenter(img, 1000, 1000), 1000, 1000, imaging.Lanczos)
	err := jpeg.Encode(&imgBuffer, thumbnail, &jpeg.Options{Quality: 90})
	if err != nil {
		return nil, err
	}
	return &imgBuffer, nil
}
