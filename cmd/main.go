package main

import (
	"github.com/adityapp/sheds/configs"
	router "github.com/adityapp/sheds/internal/apis"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetReportCaller(true)

	configs.Init()

	app, err := router.Init()
	if err != nil {
		panic(err)
	}

	app.Start()
}
