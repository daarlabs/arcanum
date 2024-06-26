package env

import (
	"log"
	"os"
)

const (
	envVar      = "APP_ENV"
	development = "development"
	staging     = "staging"
	production  = "production"
)

func EnableProductionMode() {
	if err := os.Setenv(envVar, production); err != nil {
		log.Fatalln(err)
	}
}

func EnableDevelopmentMode() {
	if err := os.Setenv(envVar, development); err != nil {
		log.Fatalln(err)
	}
}

func EnableStagingMode() {
	if err := os.Setenv(envVar, staging); err != nil {
		log.Fatalln(err)
	}
}

func Get() string {
	return os.Getenv(envVar)
}

func Development() bool {
	return Get() == development
}

func Empty() bool {
	return len(Get()) == 0
}

func Staging() bool {
	return Get() == staging
}

func Production() bool {
	return Get() == production
}
