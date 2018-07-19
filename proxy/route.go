package proxy

import (
	"net/http"
)

var controller = &Controller(Repository: Repository{})

type Route struct {
	Name		string
	Method		string
	Pattern     string
	HandlerFunc	http.HandleFunc	
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	
	for _, route := range routes { 
		var handler http.Handler
	 	handler = route.HandlerFunc
	 	handler = logger.Logger(handler, route.Name)
		   
		 router.
	  		Methods(route.Method).
	  		Path(route.Pattern).
	  		Name(route.Name).
	  		Handler(handler)
   	}
	
	return router
}