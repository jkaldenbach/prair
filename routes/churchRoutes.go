package routes

import (
	h "kaldenbach.design/prair/handlers"
)

var churchRoutes = Routes{
	Route{"ChurchIndex", "GET", "/church", h.ChurchIndex},
	Route{"ChurchCreate", "POST", "/church", h.ChurchCreate},
}
