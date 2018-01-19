package routes

import (
	h "kaldenbach.design/prayr/handlers"
)

var userRoutes = Routes{
	Route{"Index", "GET", "/", h.Index},
}
