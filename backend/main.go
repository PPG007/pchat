package main

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

var (
	httpHost = flag.String("httpHost", "0.0.0.0", "the http server listening host")
	httpPort = flag.Int("httpPort", 8080, "the http server listening port")
	wsHost   = flag.String("wsHost", "0.0.0.0", "the websocket server listening host")
	wsPort   = flag.Int("wsPort", 8081, "the websocket server listening port")
)

func main() {
	loadConfig()
}

func loadConfig() {
	mongoUri := os.Getenv("MONGO_URI")
	mongoDatabase := os.Getenv("MONGO_DATABASE")
	viper.Set("mongo", map[string]string{
		"uri":      mongoUri,
		"database": mongoDatabase,
	})
	viper.BindPFlag("httpHost", flag.Lookup("httpHost"))
	viper.BindPFlag("httpPort", flag.Lookup("httpPort"))
	viper.BindPFlag("wsHost", flag.Lookup("wsHost"))
	viper.BindPFlag("wsPort", flag.Lookup("wsPort"))
	flag.Parse()
	viper.MergeInConfig()
}
