package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	m "kaldenbach.design/prayr/models"
)

// ChurchIndex : list churches
func ChurchIndex(w http.ResponseWriter, r *http.Request) {
	var churches m.Churches
	dbErr := m.DB.Find(&churches).Error
	if dbErr != nil {
		panic(dbErr)
	}

	jsonErr := json.NewEncoder(w).Encode(churches)
	if jsonErr != nil {
		panic(jsonErr)
	}
}

// ChurchCreate : create church
func ChurchCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	church := m.Church{
		Name: name,
	}
	dbErr := m.DB.Create(&church).Error
	if dbErr != nil {
		panic(dbErr)
	}

	jsonErr := json.NewEncoder(w).Encode(church)
	if jsonErr != nil {
		panic(jsonErr)
	}
}
