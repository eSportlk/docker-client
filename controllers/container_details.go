package controllers

import (
	"../adapters"
	"../models"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type HandleContainerDetails struct {
}

func (HandleContainerDetails) ServeHTTP(w http.ResponseWriter, r *http.Request){
	cli := adapters.GetClient()
	var containerId = mux.Vars(r)["id"]
	var container, err = cli.ContainerInspect(context.Background(), containerId)
	if err!=nil{
		errResp := models.Response{
			Message: err.Error(),
			Code:    400,
		}
		response, _ := json.Marshal(errResp)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(container)
	_, _ = w.Write(response)
}