package app

import (
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"github.com/robjporter/go-xtools/xcron"
	"github.com/robjporter/go-xtools/xhealth"
	"github.com/robjporter/go-xtools/xconfig"
)

type Application struct {
	Registry  *xconfig.Config
	Logger    *logrus.Logger
	Server    *iris.Application
	LastError error
	Crons     *xcron.CronJob
	Port      int
	StartTime int64
	Checkers  *xhealth.Pinger
}
