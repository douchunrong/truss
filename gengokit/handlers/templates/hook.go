package templates

const Hook = `
package handlers

import (
	"fmt"
	"{{.ImportPath -}} /svc"
	"os"
	"os/signal"
	"syscall"
)

`

const HookInterruptHandler = `
func InterruptHandler(errc chan<- error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	terminateError := fmt.Errorf("%s", <-c)

	// Place whatever shutdown handling you want here

	errc <- terminateError
}
`
const HookSetConfig = `
func SetConfig(cfg svc.Config) svc.Config {
	// Place add custom ctx value 
	ctxMap := make(map[string]interface{})
	svc.SetCustomCtx(ctxMap)
    svc.SetRunMode("dev")
	// set CORS
    SetCORS()

	return cfg
}

func SetCORS() {
	svc.ResponseHeaderMap = make(map[string]string)
	svc.ResponseHeaderMap["Access-Control-Allow-Origin"] = "*"
	svc.ResponseHeaderMap["Access-Control-Allow-Methods"] = "POST, GET, OPTIONS, PUT, DELETE, UPDATE"
	svc.ResponseHeaderMap["Access-Control-Allow-Headers"] = "*"
	svc.ResponseHeaderMap["Access-Control-Expose-Headers"] = "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers"
	svc.ResponseHeaderMap["Access-Control-Max-Age"] = "86400"
	svc.ResponseHeaderMap["Access-Control-Allow-Credentials"] = "true"
}
`
