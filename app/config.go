package app

import (
	"runtime"
	"strconv"
)

func (a *Application) getVersionInfo() {
	a.Registry.Set("app.version", VERSION)
	a.Registry.Set("system.go.version", runtime.Version())
	a.Registry.Set("system.os.version", runtime.GOOS)
	a.Registry.Set("system.os.architecture", runtime.GOARCH)
	a.Registry.Set("system.cpu.count", strconv.Itoa(runtime.NumCPU()))
	a.logDebug(nil, "Version info populated.")
}
