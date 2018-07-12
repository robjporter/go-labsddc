package app

import (
	"github.com/kataras/iris"
	"html/template"
	"github.com/robjporter/go-xtools/xif"
	"github.com/robjporter/go-xtools/xas"
)

func (a *Application) routesHome(ctx iris.Context) {
	if DEBUG {
		ctx.ViewData("PAGE_HEADERS", NOCACHE)
	} else {
		ctx.ViewData("PAGE_HEADERS", "")
	}
	PAGE_JQUERY_EXTENSION := []template.JS{
		template.JS("$('.first.circle').circleProgress({lineCap: 'round',value: .60, fill: {color: '#ffa500'},}).on('circle-animation-progress', function (event, progress, stepValue) {$(this).find('strong').text(stepValue.toFixed(2).substr(1) * 100);});"),
		template.JS("$('.second.circle').circleProgress({lineCap: 'round',}).on('circle-animation-progress', function (event, progress) {$(this).find('strong').html(Math.round(100 * progress) + '<i>%</i>');});"),
		template.JS("$('.third.circle').circleProgress({lineCap: 'round',}).on('circle-animation-progress', function (event, progress) {$(this).find('strong').html(Math.round(100 * progress) + '<i>%</i>');});"),
		template.JS("$('#newsddc').click(function () {$('.ui.longer.sddc.modal').modal('show');});"),
		template.JS("$('#newvmm').click(function () {$('.ui.longer.vmm.modal').modal('show');});"),
		template.JS("$('#newhx').click(function () {$('.ui.longer.hx.modal').modal('show');});"),
		template.JS("$('#newvsddc').click(function () {$('.ui.longer.vsddc.modal').modal('show');});"),
		template.JS("$('.sddc.button').click(function () {$.post( \"" + a.Registry.GetString("server.path") + "createsddc\", $( '#sddcform' ).serialize()).done(function( data ){alert( 'Data Loaded: ' + data );});});"),
		template.JS("$('#sddcform').form({fields:{tenantname:'empty'}});"),
	}
	ctx.ViewData("PAGE_CSS", []string{"static/css/page.index.css"})
	ctx.ViewData("PAGE_JS", []string{"static/js/jquery.circle.js"})
	ctx.ViewData("PAGE_JQUERY", PAGE_JQUERY_EXTENSION)
	ctx.ViewData("APP_NAME", "NAME")
	ctx.ViewData("PAGE_TITLE", "TITLE")
	tmp := xas.ToInt(a.callGetAPI("sddcdeployedcount"))
	ctx.ViewData("SDDC_DEPLOYED", xif.IfEquals(tmp > 0))
	ctx.ViewData("SDDC_DEPLOYED_COUNT", tmp)
	tmp = xas.ToInt(a.callGetAPI("vmmdeployedcount"))
	ctx.ViewData("VMM_DEPLOYED", xif.IfEquals(tmp > 0))
	ctx.ViewData("VMM_DEPLOYED_COUNT", tmp)
	tmp = xas.ToInt(a.callGetAPI("hxdeployedcount"))
	ctx.ViewData("HX_DEPLOYED", xif.IfEquals(tmp > 0))
	ctx.ViewData("HX_DEPLOYED_COUNT", tmp)
	tmp = xas.ToInt(a.callGetAPI("vsddcdeployedcount"))
	ctx.ViewData("VSDDC_DEPLOYED", xif.IfEquals(tmp > 0))
	ctx.ViewData("VSDDC_DEPLOYED_COUNT", tmp)
	tmp = xas.ToInt(a.callGetAPI("vnetworkdeployedcount"))
	ctx.ViewData("VNETWORK_DEPLOYED", xif.IfEquals(tmp > 0))
	ctx.ViewData("VNETWORK_DEPLOYED_COUNT", tmp)
	tmp = xas.ToInt(a.callGetAPI("vmdeployedcount"))
	ctx.ViewData("VM_DEPLOYED", xif.IfEquals(tmp > 0))
	ctx.ViewData("VM_DEPLOYED_COUNT", tmp)
	if err := ctx.View("pages/index.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}
