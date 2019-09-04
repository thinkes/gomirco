package service

import "github.com/gorilla/mux"

func NewRoute() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Patten).
			Name(route.Name).
			Handler(route.HandleFunc)
	}

	return router
}
