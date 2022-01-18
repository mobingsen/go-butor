package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	config2 "webssh/pkg/config"
	router2 "webssh/pkg/router"
)

//1.copy the dir of ~/.kube/config to local ~
//go get github.com/gorilla/websocket
func main() {
	engine := gin.Default()
	gin.SetMode(gin.DebugMode)
	router2.InitRouter(engine)
	err := engine.Run(fmt.Sprintf("%s:%d", config2.GetString(config2.ServerHost), config2.GetInt(config2.ServerPort)))
	if err != nil {
		klog.Fatal(err)
		return
	}
}
