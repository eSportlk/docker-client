package controllers

import (
	"../adapters"
	"../models"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type HandleContainerStop struct {
}

func (HandleContainerStop) ServeHTTP(w http.ResponseWriter, r *http.Request){
	cli := adapters.GetClient()
	ctx := context.Background()
	containerId := mux.Vars(r)["id"]
	if err := cli.ContainerStop(ctx, containerId, nil); err!=nil{
		errResponse := models.Response{
			Message: err.Error(),
			Code:    models.SERVER_ERROR,
		}
		response, _ := json.Marshal(errResponse)
		w.Write(response)
		return
	}
	successResponse := models.Response{
		Message: "Container stopped successfully",
		Code:    models.SUCCESS,
	}
	response,_ := json.Marshal(successResponse);
	w.Write(response)
}