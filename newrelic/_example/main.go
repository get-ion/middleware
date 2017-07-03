package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"

	"github.com/get-ion/middleware/newrelic"
)

func main() {
	app := ion.New()
	config := newrelic.Config("APP_SERVER_NAME", "NEWRELIC_LICENSE_KEY")
	config.Enabled = true
	m, err := newrelic.New(config)
	if err != nil {
		panic(err)
	}
	app.Use(m.ServeHTTP)

	app.Get("/", func(ctx context.Context) {
		ctx.Writef("success!\n")
	})

	app.Run(ion.Addr(":8080"))
}
