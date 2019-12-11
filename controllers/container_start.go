package controllers

import (
	"../adapters"
	"../models"
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/gorilla/mux"
	"net/http"
)

type HandleContainerStart struct {
}

func (HandleContainerStart) ServeHTTP(w http.ResponseWriter, r *http.Request){
	cli := adapters.GetClient()
	ctx := context.Background()
	containerId := mux.Vars(r)["id"]
	if err := cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{}); err!=nil{
		errResponse := models.Response{
			Message: err.Error(),
			Code:    models.SERVER_ERROR,
		}
		response, _ := json.Marshal(errResponse)
		w.Write(response)
		return
	}
	successResponse := models.Response{
		Message: "Container started successfully",
		Code:    models.SUCCESS,
	}
	response,_ := json.Marshal(successResponse);
	w.Write(response)
}
