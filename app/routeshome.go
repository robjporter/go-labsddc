package app

import (
	"github.com/kataras/iris"
)

func (a *Application) routesHome(ctx iris.Context) {
	if DEBUG {
		ctx.ViewData("PAGE_HEADERS", NOCACHE)
	} else {
		ctx.ViewData("PAGE_HEADERS", "")
	}
	ctx.ViewData("PAGE_CSS", []string{"static/css/site.css"})
	ctx.ViewData("PAGE_JS", []string{"static/js/jquery-round.js"})
	ctx.ViewData("PAGE_JQUERY", "")
	ctx.ViewData("APP_NAME", "NAME")
	ctx.ViewData("PAGE_TITLE", "TITLE")
	ctx.ViewData("SDDC_DEPLOYED", false)
	ctx.ViewData("SDDC_DEPLOYED_COUNT", 4)
	ctx.ViewData("VMM_DEPLOYED", false)
	ctx.ViewData("SDDC_VMM_DEPLOYED_COUNT", 4)
	ctx.ViewData("HX_DEPLOYED", false)
	ctx.ViewData("SDDC_HX_DEPLOYED_COUNT", 4)
	ctx.ViewData("VSDDC_DEPLOYED", false)
	ctx.ViewData("VSDDC_DEPLOYED_COUNT", 4)
	ctx.ViewData("VNETWORK_DEPLOYED", false)
	ctx.ViewData("VNETWORK_DEPLOYED_COUNT", 0)
	ctx.ViewData("VM_DEPLOYED", false)
	ctx.ViewData("VM_DEPLOYED_COUNT", 0)
	if err := ctx.View("pages/index.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}
