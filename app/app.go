package app

import (
	"sync"
	"time"
	"runtime"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"github.com/robjporter/go-xtools/xcron"
	"github.com/robjporter/go-xtools/xhealth"
)

const (
	PORT = 8080
)

var (
	once     sync.Once
	instance *Application
)

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
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

	app.setupServer()
	app.setupErrorHandler()
	app.setupRoutes()
	app.setupTemplates()
	return app
}

func (a *Application) setupServer()       {}
func (a *Application) setupErrorHandler() {}
func (a *Application) setupRoutes()       {}
func (a *Application) setupTemplates()    {}
func (a *Application) Start()             {}
