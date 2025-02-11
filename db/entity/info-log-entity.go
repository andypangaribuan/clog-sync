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

type InfoLog struct {
	Id         string    `db:"id"`
	Uid        string    `db:"uid"`
	UserId     *string   `db:"user_id"`
	PartnerId  *string   `db:"partner_id"`
	Xid        *string   `db:"xid"`
	SvcName    string    `db:"svc_name"`
	SvcVersion string    `db:"svc_version"`
	SvcParent  *string   `db:"svc_parent"`
	Message    string    `db:"message"`
	Severity   string    `db:"severity"`
	Path       string    `db:"path"`
	Function   string    `db:"function"`
	Data       *string   `db:"data"`
	CreatedAt  time.Time `db:"created_at"`
}
