package app

import (
	"github.com/kataras/iris"
	"github.com/robjporter/go-xtools/xrequests"
	"time"
	"net/http"
	"errors"
	"github.com/robjporter/go-xtools/xjquery"
	"github.com/robjporter/go-xtools/xas"
	"strings"
)

var (
	token = ""
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
		tens, err := a.Registry.GetStringSlice("aci.tenants.all")
		if err != nil {
			a.api_get_tenant_all()
		}
		tens, err = a.Registry.GetStringSlice("aci.tenants.all")
		count := 0

		for i := 0; i < len(tens); i++ {
			if strings.HasPrefix(tens[i], "sddc_") {
				count++
			}
		}
		a.Registry.SetCustom("api.sddc.deployed.count", xas.ToString(count), time.Duration(15*time.Second))
		data = xas.ToString(count)
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
		tens, err := a.Registry.GetStringSlice("aci.tenants.all")
		if err != nil {
			a.api_get_tenant_all()
		}
		tens, err = a.Registry.GetStringSlice("aci.tenants.all")
		count := 0

		for i := 0; i < len(tens); i++ {
			if strings.HasPrefix(tens[i], "vsddc_") {
				count++
			}
		}
		a.Registry.SetCustom("api.vsddc.deployed.count", xas.ToString(count), time.Duration(15*time.Second))
		data = xas.ToString(count)
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
	var err error
	name := "sddc_" + ctx.PostValue("tenantname")
	if token == "" {
		token, err = getACILoginToken(a.Registry.GetString("aci.1.url"), a.Registry.GetString("aci.1.username"), a.Registry.GetString("aci.1.password"))
	}
	if err == nil {
		url := a.Registry.GetString("aci.1.url") + "api/mo/uni/tn-" + name + ".json"
		data := "{\"fvTenant\" : {\"attributes\" : {}}}"

		req := xrequests.New().SetDoNotClearBetweenSessions().SetClearSendDataBetweenSessions().IgnoreTLSCheck().SendRawBody().Timeout(10 * time.Second).Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError)
		response, _, err2 := req.SetHeader("Cookie", "APIC-Cookie="+token).Post(url).Send(data).End()
		if err2 != nil {
			ctx.Write([]byte("There has been an error when adding the SDDC."))
		}

		if response.StatusCode == 200 {
			a.Registry.Set("aci.sddc.name", name)
			ctx.Write([]byte("SDDC " + name + " has been created successfully."))
		}

		getACILogout(a.Registry.GetString("aci.1.url"), a.Registry.GetString("aci.1.username"))
	}
}

func (a *Application) api_get_tenant_all() ([]string, error) {
	if token == "" {
		token, _ = getACILoginToken(a.Registry.GetString("aci.1.url"), a.Registry.GetString("aci.1.username"), a.Registry.GetString("aci.1.password"))
	}
	url := a.Registry.GetString("aci.1.url") + "api/node/class/fvTenant.json"
	req := xrequests.New().SetDoNotClearBetweenSessions().SetClearSendDataBetweenSessions().IgnoreTLSCheck().SendRawBody().Timeout(10 * time.Second).Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError)
	response, body, err2 := req.SetHeader("Cookie", "APIC-Cookie="+token).Get(url).End()
	if response.StatusCode == 200 {
		jq := gojsonq.New().JSONString(body)
		count := jq.Reset().From("imdata").Count()
		var tenants []string
		for i := 0; i < count; i++ {
			tenants = append(tenants, xas.ToString(jq.Reset().Find("imdata.["+xas.ToString(i)+"].fvTenant.attributes.name")))
		}
		a.Registry.SetCustom("aci.tenants.all", tenants, 30*time.Second)
		return tenants, nil
	}
	return nil, err2[0]
}

func getACILogout(address, username string) error {
	var err error
	if token != "" {
		url := address + "api/aaaLogout.json"
		data := "{\"aaaUser\":{\"attributes\": {\"name\": \"" + username + "\"}}}"
		req := xrequests.New().SetDoNotClearBetweenSessions().SetClearSendDataBetweenSessions().IgnoreTLSCheck().SendRawBody().Timeout(10 * time.Second).Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError)
		response, _, err := req.SetHeader("Cookie", "APIC-Cookie="+token).Post(url).Send(data).End()
		if err == nil {
			if response.StatusCode == 200 {
				return nil
			}
		}
	}
	return err
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
		token = xas.ToString(jq.Reset().Find("imdata.[0].aaaLogin.attributes.token"))
		return token, nil
	}

	return "", errors.New("An unknown error has occured.")
}
