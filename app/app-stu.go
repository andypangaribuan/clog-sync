/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package app

import "github.com/andypangaribuan/gmod/ice"

type stuEnv struct {
	AppName        string
	AppEnv         ice.AppEnv
	AppTimezone    string
	AppAutoRecover bool
	RestPort       int

	DbSource      *stuDb
	DbDestination *stuDb
}

type stuDb struct {
	Host string
	Port int
	Name string
	User string
	Pass string
}
