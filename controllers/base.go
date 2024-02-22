package controllers

import (
	"context"
	"net/http"
)

type Base struct{}
type BaseRes struct{}

func (r *Base) Controller(context.Context) (int, string, *BaseRes, error) {
	return http.StatusOK, "Welcome to frame api", nil, nil
}
