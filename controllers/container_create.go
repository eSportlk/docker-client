package controllers

import (
	"../models"
	"encoding/json"
	"net/http"
)

type HandleContainerCreate struct {
}

func (HandleContainerCreate) ServeHTTP(w http.ResponseWriter, r *http.Request){
	var images map[string]models.Container = make(map[string]models.Container)
	response, _ := json.Marshal(images)
	_, _ = w.Write(response)
}