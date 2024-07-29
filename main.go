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

	"github.com/andypangaribuan/gmod/fm"
	"github.com/andypangaribuan/gmod/server"
)

func main() {
	fm.CallOrderedInit()
	server.FuseR(app.Env.RestPort, rest)
}

func rest(router server.RouterR) {
}
