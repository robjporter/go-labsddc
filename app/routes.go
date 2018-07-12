package app

import (
	"github.com/kataras/iris"
)

func (a *Application) setupRoutes() {
	a.Server.Get("/nolayout", func(ctx iris.Context) {
		ctx.ViewLayout(iris.NoLayout)
		if err := ctx.View("index.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	a.Server.Get("/layout", func(ctx iris.Context) {
		if err := ctx.View("index.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	a.applicationAPIRoutes()

	a.Server.Get("/", a.routesHome)
	a.Server.Get("/step1", a.routesHomeStep1)
	a.Server.Get("/step2", a.routesHomeStep2)
	a.Server.Get("/step3", a.routesHomeStep3)
	a.Server.Get("/step4", a.routesHomeStep4)
	a.Server.Get("/step5", a.routesHomeStep5)
	a.Server.Get("/step6", a.routesHomeStep6)
}

func (a *Application) applicationAPIRoutes() {
	a.Server.Get("/sddcdeployedcount", a.api_get_sddc_deployed_count)
	a.Server.Get("/vmmdeployedcount", a.api_get_vmm_deployed_count)
	a.Server.Get("/hxdeployedcount", a.api_get_hx_deployed_count)
	a.Server.Get("/vsddcdeployedcount", a.api_get_vsddc_deployed_count)
	a.Server.Get("/vnetworkdeployedcount", a.api_get_vnetwork_deployed_count)
	a.Server.Get("/vmdeployedcount", a.api_get_vm_deployed_count)
	a.Server.Post("/createsddc", a.api_post_sddc_create)
}
