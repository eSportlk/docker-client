package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"../adapters"
	"github.com/docker/docker/api/types"
)

type HandleContainerList struct {
}

func (HandleContainerList) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	cli := adapters.GetClient()
	var containersMap = make(map[string]types.Container)

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All:true})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		containersMap[container.ID[:10]] = container
	}

	response, _ := json.Marshal(containersMap)
	_, _ = w.Write(response)
}
