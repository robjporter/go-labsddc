package app

import (
	"runtime"
	"path/filepath"
	"strings"
	"github.com/sirupsen/logrus"
	"github.com/robjporter/go-xtools/xgraphics"
	"github.com/kataras/iris"
	"time"
)

func (a *Application) logDebug(fields map[string]interface{}, message string) {
	if DEBUG {
		if fields == nil {
			fields = make(map[string]interface{})
		}
		pc, file, line, _ := runtime.Caller(1)
		fn := runtime.FuncForPC(pc)

		dotName := filepath.Ext(fn.Name())
		fnName := strings.TrimLeft(dotName, ".") + "()"

		fields[".Task Number"] = count
		fields["File"] = filepath.Base(file)
		fields["Function"] = fnName
		fields["Line"] = line
		count++
		a.Logger.WithFields(logrus.Fields(fields)).Debug(message)
	}
}

func (a *Application) debug(fields map[string]interface{}, message string) {
	if DEBUG {
		if fields == nil {
			fields = make(map[string]interface{})
		}

		fields[".Task Number"] = count
		count++
		a.Logger.WithFields(logrus.Fields(fields)).Debug(message)
	}
}

func (a *Application) info(fields map[string]interface{}, message string) {
	if fields == nil {
		fields = make(map[string]interface{})
	}

	fields[".Task Number"] = count
	count++
	a.Logger.WithFields(logrus.Fields(fields)).Info(message)
}

func getMethodColor(color interface{}) string {
	g := xgraphics.New()
	tmp := color.(string)
	switch tmp {
	case "GET":
		return g.Colors.NewString().Fore(g.Colors.Color["GREEN"]).Text(tmp).Reset().String()
	case "DELETE":
		return g.Colors.NewString().Fore(g.Colors.Color["RED"]).Text(tmp).Reset().String()
	case "POST":
		return g.Colors.NewString().Fore(g.Colors.Color["BLUE"]).Text(tmp).Reset().String()
	}
	return tmp
}

func logThisMiddleware(ctx iris.Context) {
	method := getMethodColor(ctx.Method())
	ctx.Application().Logger().Infof("[SDDC] %v | %s | %s | %s | %s", time.Now().Format("02-01-2006 | 15:04:05.000000000"), method, "", ctx.Path(), ctx.RemoteAddr())
	ctx.Next()
}
