package app

import (
	"github.com/kataras/iris"
	"../data"
)

func (a *Application) SetupViews(viewsDir string) {
	a.Server.RegisterView(iris.HTML(viewsDir, ".html").Layout("shared/layout.html"))
}

func (a *Application) setupTemplates() {
	tmpl := iris.HTML("./public", ".html")
	tmpl.Layout("layout/layout.html")
	tmpl.Binary(data.Asset, data.AssetNames)
	if DEBUG {
		tmpl.Reload(true)
	}
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})
	a.Server.StaticEmbedded("/static", "./public", data.Asset, data.AssetNames)

	a.Server.RegisterView(tmpl)
}
