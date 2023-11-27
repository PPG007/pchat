package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"pchat/controller"
	"pchat/cron"
	"pchat/middleware"
)

var (
	httpHost = flag.String("httpHost", "0.0.0.0", "the http server listening host")
	httpPort = flag.Int("httpPort", 8080, "the http server listening port")
)

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
	e := gin.New()
	middleware.RegisterMiddlewares(e)
	controller.RegisterControllers(e)
	e.Run(fmt.Sprintf("%s:%d", viper.GetString("httpHost"), viper.GetInt("httpPort")))
}
