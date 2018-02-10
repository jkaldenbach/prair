package routes

import (
	h "kaldenbach.design/prair/handlers"
)

var userRoutes = Routes{
	Route{"Index", "GET", "/", h.Index},
}
