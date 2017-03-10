package route

import (
	"cf-pagerduty-service/pagerdutyapi/handle"
	"net/http"

	"github.com/gorilla/mux"
)

// Route a single API route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes all API routes
type Routes []Route

// NewRouter create a new router for the PagerDuty API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Trigger",
		"POST",
		"/pd/v1/trigger",
		handle.Trigger,
	},
}
