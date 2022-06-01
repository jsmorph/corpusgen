package corpusgen

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

func EventBridgeMatches(ctx context.Context, svc *eventbridge.Client, pattern, event string) (bool, error) {
	if svc == nil {
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return false, err
		}
		svc = eventbridge.NewFromConfig(cfg)
	}

	in := &eventbridge.TestEventPatternInput{
		Event:        aws.String(event),
		EventPattern: aws.String(pattern),
	}
	out, err := svc.TestEventPattern(ctx, in)
	if err != nil {
		log.Printf("aws error %v", err)
		return false, err
	}
	return out.Result, nil
}
