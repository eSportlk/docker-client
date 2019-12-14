package controllers

import (
	"../adapters"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"strings"
)

type HandleImagePull struct {
}

type longLatStruct struct {
	Long float64 `json:"longitude"`
	Lat  float64 `json:"latitude"`
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

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *longLatStruct)
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (HandleImagePull) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cli := adapters.GetClient()
	ctx := context.Background()
	imageName := mux.Vars(r)["image"]

	events, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
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
		fmt.Printf("EVENT: %+v\n", event)
	}

	if event != nil {
		if strings.Contains(event.Status, fmt.Sprintf("Downloaded new image of %s", imageName)){
			fmt.Println("new")
		}

		if strings.Contains(event.Status, fmt.Sprintf("Image is uptodate for %s", imageName)){
			fmt.Println("up-to-date")
		}
	}

}

func LongLatHandler(w http.ResponseWriter, r *http.Request) {
	var coodinates longLatStruct
	if err := json.NewDecoder(r.Body).Decode(&coodinates); err != nil {
		log.Printf("ERROR: %s", err)
	}
	defer r.Body.Close()
	go write(&coodinates)
	if len(clients) > 0 {
		go Echo()
	}
}

func write(coord *longLatStruct) {
	broadcast <- coord
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
		latlong := fmt.Sprintf("%f %f", val.Lat, val.Long)
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(latlong))
			if err != nil {
				log.Printf("websocket error: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
