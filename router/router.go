package router

import "github.com/gorilla/mux"
import "../controllers"

var Router = mux.NewRouter()

func GetRouter() *mux.Router {
	Router.Handle("/images", controllers.HandleImageList{})
	return Router
}
