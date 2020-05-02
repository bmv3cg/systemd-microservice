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
		ApIIndex,
	},

	Route{
		"Health",
		"GET",
		"/api/v1/service/health",
		Health,
	},
	
	Route{
		"StartService",
		"GET",
		"/api/v1/service/start",
		StartUnit,
	},
	
	Route{
		"StopService",
		"GET",
		"/api/v1/service/stop",
		StopUnit,
	},
	
	Route{
		"RestartService",
		"GET",
		"/api/v1/service/restart",
		RestartUnit,
	},
}
