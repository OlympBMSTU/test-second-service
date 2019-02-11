package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/OlympBMSTU/test-second-service/environment"
)

func DefaultErrorLog(err error, where string) {
	if err != nil {
		log.Println(fmt.Sprintf("Error occoured %s: %s", where, err.Error()))
		bytes, _ := json.Marshal(err)
		log.Println(string(bytes))
		panic(err)
	} else {
		log.Print(fmt.Sprintf("No error %s", where))
	}
}

func main() {
	// todo read config path from args
	// defautl now
	// environment
	cfgPath := "/etc/testing_service.conf"
	env, err := environment.Init(cfgPath)
	DefaultErrorLog(err, "create environment")
	err = env.Run()
	DefaultErrorLog(err, "run server")
	// router := fasthttprouter.New()
	// router.GET("/", framework_controllers.ServeSaveAnswer)
	// fasthttp.ListenAndServe("127.0.0.1:5469", router.Handler)
}
