package env

import "os"

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
