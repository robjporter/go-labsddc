package app

import (
	"time"
	"runtime"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"github.com/robjporter/go-xtools/xcron"
	"github.com/robjporter/go-xtools/xhealth"
	"github.com/robjporter/go-xtools/xconfig"
	"github.com/robjporter/go-xtools/xas"
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
	app.setupVariables()
	app.loadConfig()
	app.setupServer()
	app.SetupErrorHandlers()
	app.setupRoutes()
	app.setupTemplates()
	app.SetupWebsockets("/chat", handleConnection)

	return app
}

func (a *Application) setupVariables() {
	a.Registry.Set("server.path", "http://localhost:8080/")
}

func (a *Application) processCommandLineArguments() {
	switch kingpin.Parse() {
	case "run":
		a.start()
	case "show":
		a.showConfig()
	case "add apic":
		a.addAPIC(xas.ToString(*addAPICIP), xas.ToString(*addAPICUsername), xas.ToString(*addAPICPassword))
	case "add hx":
		a.addHX(xas.ToString(*addHXIP), xas.ToString(*addHXUsername), xas.ToString(*addHXPassword))
	case "add vc":
		a.addVC(xas.ToString(*addVCIP), xas.ToString(*addVCUsername), xas.ToString(*addVCPassword))
	case "delete apic":
		a.deleteAPIC(xas.ToString(*addAPICIP))
	case "delete hx":
		a.deleteHX(xas.ToString(*addHXIP))
	case "delete vc":
		a.deleteVC(xas.ToString(*addVCIP))
	}
}
