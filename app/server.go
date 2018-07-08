package app

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/robjporter/go-xtools/xas"
)

func (a *Application) shutdown() {
	a.Crons.Stop()
	a.info(nil, "All services shut down successfully.")
}

func (a *Application) setupServer() {
	iris.RegisterOnInterrupt(func() {
		a.info(nil, "Shutting down all services.")
		a.shutdown()
	})
	a.Server.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		// .Values are used to communicate between handlers, middleware.
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}

		ctx.Writef("(Unexpected) internal server error")
	})
	a.Server.Logger().Info("orm failed to initialized User table: %v", "TEST")
	a.Server.Logger().SetLevel("debug")
	a.Server.Use(recover.New())
	a.Server.Use(logger.New())
}

func (a *Application) Start() {
	a.logDebug(nil, "Starting services.")
	a.Crons.Run()
	a.logDebug(nil, "Cron jobs started")
	a.Server.Run(iris.Addr(":"+xas.ToString(PORT)), iris.WithConfiguration(iris.Configuration{ // default configuration:
		DisableStartupLog: true,
		DisableInterruptHandler: false,
		DisablePathCorrection: false,
		EnablePathEscape: false,
		FireMethodNotAllowed: false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode: false,
		TimeFormat: "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset: "UTF-8",
	}))
	a.logDebug(nil, "Web Server started.")
	a.info(nil, "All services have been started successfully.")
}
