package controllers

import (
	"../adapters"
	"context"
	"encoding/json"
	"net/http"
)

type HandleContainerDetails struct {
}

func (HandleContainerDetails) ServeHTTP(w http.ResponseWriter, r *http.Request){
	cli := adapters.GetClient()
	var containerId string = "abc";
	var image, err = cli.ContainerInspect(context.Background(), containerId)
	if err!=nil{
		panic(err)
	}
	response, _ := json.Marshal(image)
	_, _ = w.Write(response)
}
