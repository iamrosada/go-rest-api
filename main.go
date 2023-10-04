package main

import (
	"github.com/iamrosada/go-rest-api/config"
	"github.com/iamrosada/go-rest-api/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		//panic(err)
		// fmt.Println(err)
		logger.Errorf("Config Initialization error: %v", err)
		return
	}
	//Initialize Route
	router.Initialize()

}
