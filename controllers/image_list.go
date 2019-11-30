package controllers

import (
	"../adapters"
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"net/http"
)

type HandleImageList struct {
}

func (HandleImageList) ServeHTTP(w http.ResponseWriter, r *http.Request){
	cli := adapters.GetClient()
	var imagesMap = make(map[string]types.ImageSummary)

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	for _, image := range images {
		imagesMap[image.ID[:10]] = image
	}

	response, _ := json.Marshal(imagesMap)
	_, _ = w.Write(response)
}
