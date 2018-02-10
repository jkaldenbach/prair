package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "kaldenbach.design/prair/models"
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
	name := r.FormValue("name")

	var c m.Church
	decoder := json.NewDecoder(r.Body)
	decErr := decoder.Decode(&c)
	if decErr != nil {
		fmt.Println(decErr)
		panic(decErr)
	}
	fmt.Println(c)

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
