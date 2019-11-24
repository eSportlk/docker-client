package controllers

import (
	"../models"
	"encoding/json"
	"net/http"
)

type HandleImageList struct {
}

func (HandleImageList) ServeHTTP(w http.ResponseWriter, r *http.Request){
	var images map[string]models.Image = make(map[string]models.Image)
	response, _ := json.Marshal(images)
	_, _ = w.Write(response)
}
