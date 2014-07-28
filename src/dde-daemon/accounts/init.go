/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package accounts

import "dde-daemon"

import (
	"pkg.linuxdeepin.com/lib/dbus"
	"pkg.linuxdeepin.com/lib/log"
)

func init() {
	loader.Register(&loader.Module{"accounts", Start, Stop, true})
}

var (
	logger = log.NewLogger(ACCOUNT_DEST)
)

func Start() {
	logger.BeginTracing()
	defer logger.EndTracing()

	obj := GetManager()
	if err := dbus.InstallOnSystem(obj); err != nil {
		logger.Error("Install DBus Failed:", err)
		panic(err)
	}

	obj.updateAllUserInfo()
}

func Stop() {
	obj := GetManager()
	obj.destroy()
	dbus.UnInstallObject(obj)
}