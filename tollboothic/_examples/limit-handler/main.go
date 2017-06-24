package main

import (
	"time"

	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"

	"github.com/didip/tollbooth"
	"github.com/get-ion/middleware/tollboothic"
)

func main() {
	app := ion.New()

	// Create a limiter struct.
	limiter := tollbooth.NewLimiter(1, time.Second)

	app.Get("/", tollboothic.LimitHandler(limiter), func(ctx context.Context) {
		ctx.HTML("<b>Hello, world!</b>")
	})

	app.Run(ion.Addr(":8080"))
}

// Read more at: https://github.com/didip/tollbooth
