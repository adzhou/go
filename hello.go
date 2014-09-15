package main

import (
	"github.com/hoisie/web"
)

func hello1(ctx *web.Context, val string) {
	for k, v := range ctx.Params {
		println(k, v)
	}
}

func hello(val string) string { return "hello " + val }

func main() {
	web.Get("/(.*)", hello1)
	web.Run("0.0.0.0:9999")
}
