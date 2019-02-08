package main

import (
	"fmt"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	router := fasthttprouter.New()
	fmt.Print(router)
	// router.GET("path"

	fasthttp.ListenAndServe("localhost:8080", nil)
}
