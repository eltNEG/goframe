package controllers

import (
	"context"
	"net/http"
)

type Ping struct {
	Name  string `json:"name,omitempty" validate:"required"`
	Name2 string `json:"name2,omitempty" validate:"omitempty,oneof=e f"`
	Name3 string `json:"name3,omitempty" validate:"required"`
}
type PingRes struct {
	Name    string `json:"name,omitempty"`
	Version uint64
}

func (r *Ping) Controller(context.Context) (int, string, *PingRes, error) {
	res := &PingRes{
		Name:    r.Name,
		Version: 1,
	}

	if r.Name == "e" {
		res = nil
	}
	return http.StatusOK, "success", res, nil
}
