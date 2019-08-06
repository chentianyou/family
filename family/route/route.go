package route

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route
var serverRoutes Routes

func RegisterRoute(route Route) {
	log.Printf("register handler %s at %s", route.Name, route.Pattern)
	serverRoutes = append(serverRoutes, route)
}

func ServerRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range serverRoutes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	serverMux := http.NewServeMux()
	serverMux.Handle("/", router)

	return serverMux
}