package routes

import (
	"net/http"
)

// Route : route definition
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes : list of routes
type Routes []Route

var routes = append(
	churchRoutes,
	userRoutes...,
)
