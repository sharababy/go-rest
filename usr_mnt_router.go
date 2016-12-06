package main


import (
	"net/http"
	"httprouter"
)


type Route struct{

	Method string
	Function httprouter.Handle
	Url string
}


type Routers []Route

func NewServer() http.Handler {
		
		 server := httprouter.New()

		 for _, route := range listof {
		 	server.Handle(route.Method , route.Url , route.Function)
		 }

		 return server

}
var listof = Routers{
	Route{
		"POST",
		ReceiveJSON,
		"/",
	},
	Route{
		"PUT",
		UpdateJSON,
		"/:find_type/:find_with",
	},
	Route{
		"GET",
		FindJSON,
		"/:find_type/:find_with",
	},
	Route{
		"GET",
		DumpJSON,
		"/",
	},
	Route{
		"DELETE",
		DeleteJSON,
		"/:find_type/:find_with",
	},
}

