package app

import (
	"github.com/kataras/iris"
	"github.com/robjporter/go-xtools/xregistry"
	"github.com/sirupsen/logrus"
	"github.com/robjporter/go-xtools/xcron"
	"github.com/robjporter/go-xtools/xhealth"
)

type Application struct {
	Registry  xregsitry.Xregistry
	Logger    *logrus.Logger
	Server    *iris.Application
	LastError error
	Crons     *xcron.CronJob
	Port      int
	StartTime int64
	Checkers  *xhealth.Pinger
}
