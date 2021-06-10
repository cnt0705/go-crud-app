package controllers

import (
	"net/http"
	"todo-module/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
