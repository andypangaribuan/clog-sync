/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

/* cspell: disable-next-line */
package getstatus

import "github.com/andypangaribuan/gmod/server"

func Exec(ctx server.FuseRContext) any {
	return ctx.R200OK("healthy")
}
