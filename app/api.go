package app

import (
	"github.com/kataras/iris"
	"github.com/robjporter/go-xtools/xrequests"
	"time"
	"net/http"
	"errors"
	"github.com/robjporter/go-xtools/xjquery"
	"github.com/robjporter/go-xtools/xas"
)

func (a *Application) callGetAPI(url string) interface{} {
	req := xrequests.New().SetDoNotClearBetweenSessions().SetClearSendDataBetweenSessions().IgnoreTLSCheck().SendRawBody().Timeout(10 * time.Second).Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError)
	resp, body, err := req.Get(a.Registry.GetString("server.path") + url).End()
	if err == nil {
		if resp.StatusCode == http.StatusOK {
			return body
		}
	}
	return 0
}

func (a *Application) api_get_sddc_deployed_count(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	data := a.Registry.GetString("api.sddc.deployed.count")
	if data == "" {
		a.Registry.SetCustom("api.sddc.deployed.count", "0", time.Duration(15*time.Second))
	}
	ctx.Write([]byte(data))
}

func (a *Application) api_get_vmm_deployed_count(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	data := a.Registry.GetString("api.vmm.deployed.count")
	if data == "" {
		a.Registry.SetCustom("api.vmm.deployed.count", "0", time.Duration(15*time.Second))
	}
	ctx.Write([]byte(data))
}

func (a *Application) api_get_hx_deployed_count(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	data := a.Registry.GetString("api.hx.deployed.count")
	if data == "" {
		a.Registry.SetCustom("api.hx.deployed.count", "0", time.Duration(15*time.Second))
	}
	ctx.Write([]byte(data))
}

func (a *Application) api_get_vsddc_deployed_count(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	data := a.Registry.GetString("api.vsddc.deployed.count")
	if data == "" {
		a.Registry.SetCustom("api.vsddc.deployed.count", "0", time.Duration(15*time.Second))
	}
	ctx.Write([]byte(data))
}

func (a *Application) api_get_vnetwork_deployed_count(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	data := a.Registry.GetString("api.vnetwork.deployed.count")
	if data == "" {
		a.Registry.SetCustom("api.vnetwork.deployed.count", "0", time.Duration(15*time.Second))
	}
	ctx.Write([]byte(data))
}

func (a *Application) api_get_vm_deployed_count(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	data := a.Registry.GetString("api.vm.deployed.count")
	if data == "" {
		a.Registry.SetCustom("api.vm.deployed.count", "0", time.Duration(15*time.Second))
	}
	ctx.Write([]byte(data))
}

func (a *Application) api_post_sddc_create(ctx iris.Context) {

	name := ctx.PostValue("tenantname")
	token, err := getACILoginToken(a.Registry.GetString("aci.1.url"), a.Registry.GetString("aci.1.username"), a.Registry.GetString("aci.1.password"))

	if err == nil {
		url := a.Registry.GetString("aci.1.url") + "api/mo/uni/tn-" + name + ".json"
		data := "{\"fvTenant\" : {\"attributes\" : {}}}"

		req := xrequests.New().SetDoNotClearBetweenSessions().SetClearSendDataBetweenSessions().IgnoreTLSCheck().SendRawBody().Timeout(10 * time.Second).Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError)
		response, _, err2 := req.SetHeader("Cookie", "APIC-Cookie="+token).Post(url).Send(data).End()
		if err2 != nil {
			ctx.Write([]byte("There has been an error when adding the SDDC."))
		}

		if response.StatusCode == 200 {
			ctx.Write([]byte("SDDC has been created successfully."))
		}

		getACILogout(a.Registry.GetString("aci.1.url"), token, a.Registry.GetString("aci.1.username"))
	}
}

func getACILogout(address, token, username string) error {
	url := address + "api/aaaLogout.json"
	data := "{\"aaaUser\":{\"attributes\": {\"name\": \"" + username + "\"}}}"
	req := xrequests.New().SetDoNotClearBetweenSessions().SetClearSendDataBetweenSessions().IgnoreTLSCheck().SendRawBody().Timeout(10 * time.Second).Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError)
	response, _, err := req.SetHeader("Cookie", "APIC-Cookie="+token).Post(url).Send(data).End()
	if err == nil {
		if response.StatusCode == 200 {
			return nil
		}
	}
	return err[0]
}

func getACILoginToken(address, username, password string) (string, error) {
	url := address + "api/aaaLogin.json"
	// data := "{\"username\":\"" + username + "\",\"password\":\"" + password + "\"}"
	data := "{\"aaaUser\":{\"attributes\":{\"name\":\"" + username + "\",\"pwd\":\"" + password + "\"}}}"
	req := xrequests.New().SetDoNotClearBetweenSessions().SetClearSendDataBetweenSessions().IgnoreTLSCheck().SendRawBody().Timeout(10 * time.Second).Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError)
	response, body, err := req.Post(url).Send(data).End()
	if err != nil {
		return "", err[0]
	}
	if response.StatusCode == 200 {
		jq := gojsonq.New().JSONString(body)
		return xas.ToString(jq.Reset().Find("imdata.[0].aaaLogin.attributes.token")), nil
	}

	return "", errors.New("An unknown error has occured.")
}
