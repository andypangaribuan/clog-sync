/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package main

import (
	"clog-sync/app"
	"clog-sync/handler"

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/server"
)

func main() {
	fm.CallOrderedInit()
	server.FuseR(app.Env.RestPort, rest)
}

func rest(router server.RouterR) {
	router.AutoRecover(app.Env.AppAutoRecover)
	router.PrintOnError(app.Env.AppServerPrintOnError)

	router.Endpoints(nil, nil, map[string][]func(server.FuseRContext) any{
		"GET: /private/status": {handler.Private.Status},
	})
}
