package app

import (
	"time"
	"runtime"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"github.com/robjporter/go-xtools/xcron"
	"github.com/robjporter/go-xtools/xhealth"
	"github.com/robjporter/go-xtools/xconfig"
)

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
			Registry:  xconfig.New(),
			Logger:    logrus.New(),
			Server:    iris.New(),
			Crons:     xcron.New(),
			Port:      PORT,
			StartTime: time.Now().Unix(),
			Checkers:  xhealth.New(),
		}
	})
	return instance
}

func New() *Application {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app := GetInstance()
	app.displayBanner()
	app.info(nil, "Starting application configuration.")
	app.Logger.SetLevel(logrus.DebugLevel)

	app.getVersionInfo()
	app.setupServer()
	app.SetupErrorHandlers()
	app.setupRoutes()
	app.setupTemplates()
	app.SetupWebsockets("/chat", handleConnection)

	return app
}
