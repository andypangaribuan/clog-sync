/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

/* spell-checker: disable */
package private

import (
	getstatus "clog-sync/handler/private/get-status"

	"github.com/andypangaribuan/gmod/server"
)

type Handler struct{}

func (*Handler) Status(ctx server.FuseRContext) any {
	return getstatus.Exec(ctx)
}
