package routes

import (
	h "kaldenbach.design/prayr/handlers"
)

var churchRoutes = Routes{
	Route{"ChurchIndex", "GET", "/church", h.ChurchIndex},
	Route{"ChurchCreate", "GET", "/church/{name}", h.ChurchCreate},
}
