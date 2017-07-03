package main

import (
	"time"

	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	cw "github.com/get-ion/middleware/cloudwatch"
)

func main() {
	app := ion.New()
	app.Use(cw.New("us-east-1", "test").ServeHTTP)

	app.Get("/", func(ctx context.Context) {
		put := cw.GetPutFunc(ctx)

		put([]*cloudwatch.MetricDatum{
			{
				MetricName: aws.String("MyMetric"),
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("ThingOne"),
						Value: aws.String("something"),
					},
					{
						Name:  aws.String("ThingTwo"),
						Value: aws.String("other"),
					},
				},
				Timestamp: aws.Time(time.Now()),
				Unit:      aws.String("Count"),
				Value:     aws.Float64(42),
			},
		})

		ctx.StatusCode(ion.StatusOK)
		ctx.Text("success!\n")
	})

	// http://localhost:8080
	// should give: NoCredentialProviders
	// which is correct, you have to authorize your aws, we asumme that you know how to.
	app.Run(ion.Addr(":8080"))
}
