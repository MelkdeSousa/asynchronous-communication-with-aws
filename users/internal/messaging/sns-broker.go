package messaging

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/melkdesousa/wottva/users/pkg/contracts"
)

type SNSBroker struct {
	snsClient *sns.Client
}

func NewSNSBroker() *SNSBroker {
	var (
		accessKey   = os.Getenv("AWS_ACCESS_KEY")
		secretKey   = os.Getenv("AWS_SECRET_KEY")
		region      = os.Getenv("AWS_REGION")
		snsEndpoint = os.Getenv("SNS_ENDPOINT")
	)

	options := sns.Options{
		Region:      region,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	}

	client := sns.New(options, func(o *sns.Options) {
		o.Region = region
		o.BaseEndpoint = aws.String(snsEndpoint)
	})

	return &SNSBroker{
		snsClient: client,
	}
}

func (b *SNSBroker) Publish(topic contracts.Topic, message []byte) error {
	listTopicsParams := &sns.ListTopicsInput{}

	paginator := sns.NewListTopicsPaginator(b.snsClient, listTopicsParams)

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())

		if err != nil {
			err = fmt.Errorf("failed to list topics: %v", err)
			return err
		}

		for _, t := range output.Topics {
			topicArn := *t.TopicArn

			if strings.Contains(topicArn, topic.String()) {
				_, err := b.snsClient.Publish(context.TODO(), &sns.PublishInput{
					TopicArn: aws.String(topicArn),
					Message:  aws.String(string(message)),
				})

				if err != nil {
					err = fmt.Errorf("failed to publish message: %v", err)

					return err
				}

				return nil
			}
		}
	}

	return nil
}
