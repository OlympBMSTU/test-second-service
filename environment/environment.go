package environment

import (
	"log"

	"github.com/valyala/fasthttp"

	"github.com/buaazp/fasthttprouter"

	"github.com/OlympBMSTU/exercises/logger"
	"github.com/OlympBMSTU/test-second-service/config"
	"github.com/OlympBMSTU/test-second-service/framework_controllers"
)

type ILogger logger.ILogger
type Config config.Config

// import "github.com//config"

type ServerEnvironment struct {
	config               *config.Config
	logger               ILogger
	frameworkControllers framework_controllers.FrameworkController
}

//
func Init(configPath string) (*ServerEnvironment, error) {
	// logger := logger.New()
	conf, err := config.Read(configPath)
	if err != nil {
		log.Println("Error read config")
		return nil, err
	}
	// config := Config{}
	// config.
	logger, err := logger.New(conf.GetLoggerType(), conf.GetLoggerPath(), conf.GetLoggerLevel())

	if err != nil {
		log.Println("Error create logger")
		return nil, err
	}

	frameworkControllers := framework_controllers.Init(conf.GetListenerHost(), conf.GetListenerPort())

	env := ServerEnvironment{
		config:               conf,
		logger:               logger,
		frameworkControllers: frameworkControllers,
	}

	if err != nil {
		return nil, err
	}

	err = frameworkControllers.Run()
	if err != nil {
		return nil, err
	}
	// router := fasthttprouter.New()

	return &env, nil
}

func (self ServerEnvironment) Run() error {
	router := fasthttprouter.New()
	router.GET("/", self.rawControllers.ServeSaveAnswer)
	router.GET("/api/get/", self.rawControllers.ServeSaveAnswer) //self.rawControllers.ServeSaveAnswer)
	server := ":" + self.config.GetListenerPort()
	log.Print(server)
	err := fasthttp.ListenAndServe(server, router.Handler)
	return err
}
