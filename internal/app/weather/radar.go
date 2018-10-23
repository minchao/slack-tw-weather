package weather

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/spf13/cobra"
)

var (
	radarCmd = &cobra.Command{
		Use:   "radar",
		Short: "Weather radar (Composite reflectivity)",
		RunE:  radarFunc,
	}
)

func radarFunc(_ *cobra.Command, _ []string) error {
	sess := session.Must(session.NewSession())
	svc := sns.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	_, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(responseUrl),
		TopicArn: aws.String(os.Getenv("RADAR_TOPIC_SNS_ARN")),
	})
	if err != nil {
		return err
	}
	return nil
}
