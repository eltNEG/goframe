package database

import "goframe/enum"

type permission string

var Permissions = enum.MakeEnum[permission](
	struct {
		USR_READ  permission
		USR_WRITE permission
		ADMIN     permission
	}{
		USR_READ:  "usr_read",
		USR_WRITE: "usr_write",
		ADMIN:     "admin",
	},
)
