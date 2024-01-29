package main

import (
	"fmt"
	"github.com/spf13/cast"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"pchat/controller"
	"pchat/cron"
	"pchat/middleware"
)

var (
	httpHost = flag.String("httpHost", "0.0.0.0", "the http server listening host")
	httpPort = flag.Int("httpPort", 8080, "the http server listening port")

	isDebug = cast.ToBool(os.Getenv("IS_DEBUG"))
)

// @title						PChat API
// @version					1.0
// @description				PChat 的接口文档
// @BasePath					/v1
// @securityDefinitions.apiKey	token
// @in							header
// @name						X-Access-Token
// @description				jwt string
func main() {
	loadConfig()
	InitDefaultResources()
	cron.Start()
	startGin()
}

func loadConfig() {
	viper.BindPFlag("httpHost", flag.Lookup("httpHost"))
	viper.BindPFlag("httpPort", flag.Lookup("httpPort"))
	flag.Parse()
	viper.MergeInConfig()
}

func startGin() {
	e := controller.GetRoot(isDebug)
	middleware.RegisterMiddlewares(e)
	err := e.Run(fmt.Sprintf("%s:%d", viper.GetString("httpHost"), viper.GetInt("httpPort")))
	if err != nil {
		panic(err)
	}
}
