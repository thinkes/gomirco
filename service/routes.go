package service

import "net/http"

type Route struct {
	Name       string
	Method     string
	Patten     string
	HandleFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetAccount",
		"GET",
		"/account/{accountId}",
		GetAccount,
	},
}
