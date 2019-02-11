package framework_controllers

import (
	"fmt"
	"log"

	"github.com/OlympBMSTU/test-second-service/config"
	"github.com/OlympBMSTU/test-second-service/controllers"
	"github.com/OlympBMSTU/test-second-service/logger"
	"github.com/buaazp/fasthttprouter"

	"github.com/valyala/fasthttp"
)

// type Environment environment.ServerEnvironment

type FrameworkController struct {
	controllerLayer controllers.IControllers
	router          *fasthttprouter.Router
	config          *config.Config
	loger           *logger.ILoger
	port            string
	host            string
}

// func  initPaths(router *fasthttprouter.Router, contr FrameworkController) {
// router.GET("/",
// )
// }

func Init(host, port string) FrameworkController {
	router := fasthttprouter.New()
	// initPaths(router)
	log.Print("paths inited")
	return FrameworkController{
		router: router,
		port:   port,
		host:   host,
	}
}

func (self FrameworkController) Run() error {
	server := ":" + self.port
	return fasthttp.ListenAndServe(server, self.router.Handler)
}

func (self FrameworkController) ServeSaveAnswer(ctx *fasthttp.RequestCtx) {
	fmt.Print("request")
	fmt.Fprintf(ctx, "hi")
}

// func ServeSaveAnswer(ctx *fasthttp.RequestCtx) {
// 	fmt.Print("request")
// 	fmt.Fprintf(ctx, "hi")
// }
