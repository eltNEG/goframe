package main

import (
	ctrl "goframe/controllers"
	"goframe/utils"
	"log/slog"
)

func main() {

	utils.SetLogLevel(slog.LevelInfo)

	http := utils.NewHttp()

	http.Get("/", utils.Handlerize(&ctrl.Base{}, &ctrl.BaseRes{}))

	http.Get("/ping", utils.Handlerize(&ctrl.Ping{}, &ctrl.PingRes{}))
	http.Post("/ping", utils.Handlerize(&ctrl.Ping{}, &ctrl.PingRes{}))

	http.Get("/subscribe", utils.Handlerize(&ctrl.Subscribe2{}, &ctrl.SubscribeRes{}))
	http.Get("/unsubscribe", utils.Handlerize(&ctrl.Unsubscribe{}, &ctrl.SubscribeRes{}))

	http.Get("/user", utils.Handlerize(&ctrl.User{}, &ctrl.UserRes{}))

	slog.With("key", "value").Debug("hello world")
	slog.Info("listening on http://localhost:8080")

	http.ListenAndServe(":8080")
}
