package controllers

import (
	"models"
	"net/http"
)

var HandleSocketComs = func (w http.ResponseWriter, r *http.Request) {
	hub := models.NewHub()
	go hub.Run()
	models.ServeInterviewWS(hub, w, r)
}


