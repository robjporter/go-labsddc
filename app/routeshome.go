package app

import (
	"github.com/kataras/iris"
)

func (a *Application) routesHome(ctx iris.Context) {
	if err := ctx.View("pages/index.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}
