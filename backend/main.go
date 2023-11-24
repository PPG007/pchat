package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"pchat/controller"
	"pchat/middleware"
)

var (
	httpHost = flag.String("httpHost", "0.0.0.0", "the http server listening host")
	httpPort = flag.Int("httpPort", 8080, "the http server listening port")
	wsHost   = flag.String("wsHost", "0.0.0.0", "the websocket server listening host")
	wsPort   = flag.Int("wsPort", 8081, "the websocket server listening port")
)

func main() {
	loadConfig()
	InitDefaultResources()
	startGin()
}

func loadConfig() {
	viper.BindPFlag("httpHost", flag.Lookup("httpHost"))
	viper.BindPFlag("httpPort", flag.Lookup("httpPort"))
	viper.BindPFlag("wsHost", flag.Lookup("wsHost"))
	viper.BindPFlag("wsPort", flag.Lookup("wsPort"))
	flag.Parse()
	viper.MergeInConfig()
}

func startGin() {
	e := gin.New()
	middleware.RegisterMiddlewares(e)
	controller.RegisterControllers(e)
	e.Run(fmt.Sprintf("%s:%d", viper.GetString("httpHost"), viper.GetInt("httpPort")))
}
