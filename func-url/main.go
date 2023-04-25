package main

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome!")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
}

func panicHandler(ctx *fasthttp.RequestCtx, p interface{}) {
	fmt.Println("Recovered in panicHandler", p, string(debug.Stack()))

	ctx.Response.Reset()
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
}

func main() {
	r := router.New()
	r.PanicHandler = panicHandler

	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)

	log.Fatal(fasthttp.ListenAndServe(":443", r.Handler))
}
