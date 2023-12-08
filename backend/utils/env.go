package utils

import "os"

func AppName() string {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "pchat"
	}
	return appName
}
