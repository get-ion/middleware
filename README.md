This repository provides a way to share any minor handlers for [ion](https://github.com/get-ion/ion) web framework. You can view the built'n supported handlers by pressing [here](https://github.com/get-ion/ion/tree/master/middleware).

[![Build status](https://api.travis-ci.org/get-ion/middleware.svg?branch=master&style=flat-square)](https://travis-ci.org/get-ion/middleware)

## Installation

```sh
$ go get github.com/get-ion/middleware/...
```

Middleware is just a chain handlers which can be executed before or after the main handler, can transfer data between handlers and communicate with third-party libraries, they are just functions.

### How can I register middleware?


**To a single route**
```go
app := ion.New()
app.Get("/mypath", myMiddleware1, myMiddleware2, func(ctx context.Context){}, func(ctx context.Context){}, myMiddleware5,myMainHandlerLast)
```

**To a party of routes or subdomain**
```go

myparty := app.Party("/myparty", myMiddleware1,func(ctx context.Context){},myMiddleware3)
{
	//....
}

```

**To all routes**
```go
app.Use(func(ctx context.Context){}, myMiddleware2)
```

**To global, all routes on all subdomains on all parties**
```go
app.UseGlobal(func(ctx context.Context){}, myMiddleware2)
```


## Can I use standard net/http handler with ion?

**Yes** you can, just pass the Handler inside the `handlerconv.FromStd` in order to be converted into ion.HandlerFunc and register it as you saw before.

### Convert handler which has the form of `http.Handler/HandlerFunc`

```go
package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
	"github.com/get-ion/ion/core/handlerconv"
)

func main() {
	app := ion.New()

	sillyHTTPHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	     println(r.RequestURI)
	})

	sillyConvertedToIon := handlerconv.FromStd(sillyHTTPHandler)
	// FromStd can take (http.ResponseWriter, *http.Request, next http.Handler) too!
	app.Use(sillyConvertedToIon)

	app.Run(ion.Addr(":8080"))
}

```

## Contributing

If you are interested in contributing to this project, please push a PR.

## People

[List of all contributors](https://github.com/get-ion/middleware/graphs/contributors)