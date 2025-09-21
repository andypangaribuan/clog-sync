/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package entity

import "time"

type InternalSyncLog struct {
	TableName string    `db:"table_name"`
	LastSync  time.Time `db:"last_sync"`
}
