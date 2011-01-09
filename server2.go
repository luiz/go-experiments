package main

import(
	"web"
	"fmt"
	"mustache"
)

func hello(ctx *web.Context, name string) string {
	fmt.Printf("Parameters in request:\n")
	for key, val := range ctx.Request.Params {
		fmt.Printf("%s = %s\n", key, val)
	}
	return mustache.RenderFile("hello.mustache", map[string]string{ "name": name })
}

func main() {
	web.Get("/(.+)/hello", hello)
	web.Run("127.0.0.1:9999")
}
