package utils

import (
	"context"
	"goframe/enum"
	"net/http"
)

type contextKey string

var ContextKeys = enum.MakeEnum[contextKey](struct {
	req contextKey
	res contextKey
}{
	req: "req",
	res: "res",
})

func setcontextWR(ctx context.Context, req *http.Request, res http.ResponseWriter) context.Context {
	ctx = context.WithValue(ctx, ContextKeys.V.req, req)
	ctx = context.WithValue(ctx, ContextKeys.V.res, res)
	return ctx
}

func GetOriginalReq(ctx context.Context) *http.Request {
	return ctx.Value(ContextKeys.V.req).(*http.Request)
}

func GetOriginalRes(ctx context.Context) http.ResponseWriter {
	return ctx.Value(ContextKeys.V.res).(http.ResponseWriter)
}
