package env

import (
	"os"
	"pchat/repository/bson"
)

var ServerId = ""

func init() {
	ServerId = bson.NewObjectId().Hex()
}

func GetAppName() string {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "pchat"
	}
	return appName
}

func IsDebug() bool {
	return os.Getenv("ENV") == "dev"
}
