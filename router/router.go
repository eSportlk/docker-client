package router

import "github.com/gorilla/mux"
import "../controllers"

var Router = mux.NewRouter()

func GetRouter() *mux.Router {
	//Images
	Router.Handle("/images", controllers.HandleImageList{})
	Router.Handle("/images/pull", controllers.HandleImagePull{})
	Router.Handle("/images/{id}", controllers.HandleImageDetails{})

	//Containers
	Router.Handle("/containers", controllers.HandleContainerList{})
	Router.Handle("/containers/{id}", controllers.HandleContainerDetails{})
	Router.Handle("/containers/create", controllers.HandleContainerCreate{})
	Router.Handle("/containers/{id}/start", controllers.HandleContainerStart{})
	Router.Handle("/containers/{id}/stop", controllers.HandleContainerStop{})
	Router.Handle("/containers/{id}/kill", controllers.HandleContainerKill{})
	Router.Handle("/containers/{id}/pause", controllers.HandleContainerPause{})

	//websocket
	Router.HandleFunc("/longlat", controllers.LongLatHandler).Methods("POST")
	Router.HandleFunc("/ws", controllers.WsHandler)
	return Router
}
