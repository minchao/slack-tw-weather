package radar

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(_ context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
	}
}
