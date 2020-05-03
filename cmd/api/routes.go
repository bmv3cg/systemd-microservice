package api

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"Health",
		"GET",
		"/api/v1/service/health",
		Health,
	},

	Route{
		"StartService",
		"POST",
		"/api/v1/service/start",
		StartUnit,
	},

	Route{
		"StopService",
		"POST",
		"/api/v1/service/stop",
		StopUnit,
	},

	Route{
		"RestartService",
		"POST",
		"/api/v1/service/restart",
		RestartUnit,
	},
}
