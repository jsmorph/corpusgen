package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	gen "github.com/jsmorph/corpusgen"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		defaultEvent = `{"version":"0","id":"6a7e8feb-b491-4cf7-a9f1-bf3703467718","detail-type":"EC2 Instance State-change Notification","source":"aws.ec2","account":"111122223333","time":"2017-12-22T18:43:48Z","region":"us-west-1","resources":["arn:aws:ec2:us-west-1:123456789012:instance/i-1234567890abcdef0"],"detail":{"instance-id":"i-1234567890abcdef0","state":"terminated"}}`

		defaultPattern = `{"source": ["aws.ec2"],"detail-type": ["EC2 Instance State-change Notification"], "detail": {"state": ["terminated"]}}`

		p  = flag.String("p", defaultPattern, "pattern to check")
		e  = flag.String("e", defaultEvent, "optional event for matching")
		eb = flag.Bool("eb", false, "also check with EventBridge")
	)

	flag.Parse()

	lp := gen.CheckPattern(ctx, *p, *e, *eb)

	js, err := json.Marshal(lp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", js)

}
