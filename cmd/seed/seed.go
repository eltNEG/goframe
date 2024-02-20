package main

import (
	"context"
	"goframe/database"
)

var permissions = []string{
	"usr_read",
	"usr_write",
	"admin",
}

func main() {
	// create permissions
	db, err := database.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.CreatePermissions(context.Background(), permissions...)
	if err != nil {
		panic(err)
	}

	// create user
	err = db.CreateUser(context.Background(), "admin", database.Permissions.V.ADMIN)
	if err != nil {
		panic(err)
	}
}
