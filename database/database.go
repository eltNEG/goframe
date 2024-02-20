package database

import (
	"context"
	"fmt"
	ent "goframe/ent/generated"
	"goframe/ent/generated/permissions"
	"goframe/utils"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*ent.Client
}

func GetDB() (*DB, error) {
	client, err := ent.Open("sqlite3", ".data/frame.db?_fk=1")
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}
	return &DB{client}, nil
}

func (db *DB) CreatePermissions(ctx context.Context, permission ...string) error {
	if len(permission) == 0 {
		return fmt.Errorf("no permissions to create")
	}

	if len(permission) == 1 {
		_, err := db.Client.Permissions.Create().
			SetName(permission[0]).
			Save(ctx)
		return err
	}
	perms := make([]*ent.PermissionsCreate, len(permission))
	for i, p := range permission {
		perms[i] = db.Client.Permissions.Create().SetName(p)
	}
	pc := db.Client.Permissions.CreateBulk(
		perms...,
	)
	_, err := pc.Save(ctx)
	return err
}

func (db *DB) CreateUser(ctx context.Context, username string, permission permission) error {
	gid := utils.GenID(utils.GenIDArg.WithPrefix("usr"))
	permissionEnt, err := db.Permissions.Query().Where(permissions.Name(string(permission))).Only(ctx)
	if err != nil {
		return err
	}
	_, err = db.Client.User.Create().SetUsername(username).SetGid(gid).AddPermissions(permissionEnt).Save(ctx)
	return err
}
