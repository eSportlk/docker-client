package controllers

import (
	"../models"
	"encoding/json"
	"net/http"
)

type HandleImageDetails struct {
}

func (HandleImageDetails) ServeHTTP(w http.ResponseWriter, r *http.Request){
	var images map[string]models.Image = make(map[string]models.Image)
	response, _ := json.Marshal(images)
	_, _ = w.Write(response)
}
