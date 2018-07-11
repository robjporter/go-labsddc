package app

import (
	"runtime"
	"strconv"
	"fmt"
)

func (a *Application) getVersionInfo() {
	a.Registry.Set("app.version", VERSION)
	a.Registry.Set("system.go.version", runtime.Version())
	a.Registry.Set("system.os.version", runtime.GOOS)
	a.Registry.Set("system.os.architecture", runtime.GOARCH)
	a.Registry.Set("system.cpu.count", strconv.Itoa(runtime.NumCPU()))
	a.logDebug(nil, "Version info populated.")
}

func (a *Application) showConfig() {
	fmt.Println("SHOWING CONFIG")
}

func (a *Application) addAPIC(ip, username, password string) {

}

func (a *Application) addHX(ip, username, password string) {

}

func (a *Application) addVC(ip, username, password string) {

}

func (a *Application) deleteAPIC(ip string) {

}

func (a *Application) deleteHX(ip string) {

}

func (a *Application) deleteVC(ip string) {

}
