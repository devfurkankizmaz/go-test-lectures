package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

func Test_application_routes(t *testing.T) {
	var registered = []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/static/*", "GET"},
	}
	var app application
	mux := app.routes()

	chiRoutes := mux.(chi.Routes)
	for _, route := range registered {
		// Check to see if route exists
		if !routeExists(route.route, route.method, chiRoutes) {
			t.Errorf("Route %s does not exist", route.route)
		}
	}
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	var found = false
	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == testRoute && method == testMethod {
			found = true
		}
		return nil
	})
	return found
}
