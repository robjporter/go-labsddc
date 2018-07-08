package app

import (
	"github.com/kataras/iris"
)

func (a *Application) SetupErrorHandlers() {
	a.Server.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     "APPNAME",
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}
