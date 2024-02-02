package main

import (
	ctrl "goframe/controllers"
	"goframe/utils"
	"log/slog"
)

func main() {

	utils.SetLogLevel(slog.LevelInfo)

	http := utils.NewHttp()

	http.Get("/ping", utils.Handlerize(&ctrl.Ping{}, &ctrl.PingRes{}))
	http.Post("/ping", utils.Handlerize(&ctrl.Ping{}, &ctrl.PingRes{}))

	slog.With("key", "value").Debug("hello world")
	slog.Info("listening on http://localhost:8080")

	http.ListenAndServe("localhost:8080")
}