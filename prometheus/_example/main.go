package main

import (
	"math/rand"
	"time"

	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"

	prometheusMiddleware "github.com/get-ion/middleware/prometheus"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	app := ion.New()
	m := prometheusMiddleware.New("serviceName", 300, 1200, 5000)

	app.Use(m.ServeHTTP)

	app.OnStatusCode(ion.StatusNotFound, func(ctx context.Context) {
		// error code handlers are not sharing the same middleware as other routes, so we have
		// to call them inside their body.
		m.ServeHTTP(ctx)

		ctx.Writef("Not Found")
	})

	app.Get("/", func(ctx context.Context) {
		sleep := rand.Intn(4999) + 1
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		ctx.Writef("Slept for %d milliseconds", sleep)
	})

	app.Get("/metrics", ion.ToHandler(prometheus.Handler()))

	// http://localhost:8080/
	// http://localhost:8080/anotfound
	// http://localhost:8080/metrics
	app.Run(ion.Addr(":8080"))
}
