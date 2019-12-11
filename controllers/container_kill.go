package controllers

import (
	"../adapters"
	"../models"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type HandleContainerKill struct {
}

func (HandleContainerKill) ServeHTTP(w http.ResponseWriter, r *http.Request){
	cli := adapters.GetClient()
	ctx := context.Background()
	containerId := mux.Vars(r)["id"]
	if err := cli.ContainerKill(ctx, containerId, ""); err!=nil{
		errResponse := models.Response{
			Message: err.Error(),
			Code:    models.SERVER_ERROR,
		}
		response, _ := json.Marshal(errResponse)
		w.Write(response)
		return
	}
	successResponse := models.Response{
		Message: "Container killed successfully",
		Code:    models.SUCCESS,
	}
	response,_ := json.Marshal(successResponse);
	w.Write(response)
}