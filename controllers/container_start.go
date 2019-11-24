package controllers

import (
	"../models"
	"encoding/json"
	"net/http"
)

type HandleContainerStart struct {
}

func (HandleContainerStart) ServeHTTP(w http.ResponseWriter, r *http.Request){
	var images map[string]models.Container = make(map[string]models.Container)
	response, _ := json.Marshal(images)
	_, _ = w.Write(response)
}
