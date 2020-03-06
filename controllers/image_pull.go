package controllers

import (
	"../adapters"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"strings"
)

type HandleImagePull struct {
}

type Event struct {
	Status          string `json:"status"`
	Error           string `json:"error"`
	Progress        string `json:"progress"`
	ProgressDetails struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"progressDetails"`
}

type ImageData struct {
	Name string `json:"name"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *Event)
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var echoValue = false

func (HandleImagePull) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cli := adapters.GetClient()
	ctx := context.Background()
	decoder := json.NewDecoder(r.Body)
	var imageData ImageData
	err := decoder.Decode(&imageData)

	if err != nil {
		panic(err)
	}

	if !echoValue {
		go Echo()
		echoValue = true
	}

	events, err := cli.ImagePull(ctx, imageData.Name, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(events)

	var event *Event

	for {
		if err := d.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		var eventData Event

		if err := json.NewDecoder(r.Body).Decode(&eventData); err != nil {
			log.Printf("ERROR: %s", err)
		}
		go write(event)
		fmt.Printf("EVENT: %+v\n", event)
	}

	if event != nil {
		if strings.Contains(event.Status, fmt.Sprintf("Downloaded new image of %s", imageData.Name)) {
			fmt.Println("new")
		}

		if strings.Contains(event.Status, fmt.Sprintf("Image is uptodate for %s", imageData.Name)) {
			fmt.Println("up-to-date")
		}
	}

}

func write(event *Event) {
	broadcast <- event
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	clients[ws] = true
}

func Echo() {
	for {
		val := <-broadcast
		event := fmt.Sprintf("%s %s %d %d %s", val.Status, val.Progress, val.ProgressDetails.Current, val.ProgressDetails.Total, val.Error)
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(event))
			if err != nil {
				log.Printf("websocket error: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
