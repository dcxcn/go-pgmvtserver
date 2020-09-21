package main

import (
	"github.com/kataras/iris/v12"
	"go-pgmvtserver/app/middleware"
	"go-pgmvtserver/app/router"
	"go-pgmvtserver/config"
)

func main() {
	irisApp := iris.New()
	irisApp.AllowMethods(iris.MethodOptions)
	irisApp.Use(middleware.CORS)
	router.Router(irisApp)
	configInfo := config.ConfigInfo
	irisApp.Listen(configInfo.Port)
}
