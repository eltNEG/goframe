package controllers

import (
	"context"
	"goframe/database"
	"goframe/ent/generated/user"
	"log/slog"
)

type User struct {
	Username string `json:"username,omitempty" validate:"required"`
}

type UserRes struct {
	ID          string   `json:"id,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

func (s *User) Controller(ctx context.Context) (int, string, *UserRes, error) {
	db, err := database.GetDB()
	if err != nil {
		return 500, "error getting db", nil, err
	}
	defer db.Close()
	user, err := db.User.Query().Where(user.Username(s.Username)).WithPermissions().Only(ctx)
	if err != nil {
		slog.With("username", s.Username).With("err", err).Error("error getting user")
		return 404, "user not found", nil, nil
	}
	perms := make([]string, len(user.Edges.Permissions))
	for i, p := range user.Edges.Permissions {
		perms[i] = p.Name
	}
	return 200, "ok", &UserRes{ID: user.Gid, Permissions: perms}, nil
}
