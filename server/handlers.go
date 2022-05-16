package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"schoolXBackend/datastore"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API works")
}

func cityHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		cityGetHandler(w, r)
	case "POST":
		cityAddHandler(w, r)
	case "PUT":
		cityEditHandler(w, r)
	case "DELETE":
		cityDeleteHandler(w, r)
	case "OPTIONS":
		optionsHandler(w, r)
	}
}
func cityGetHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		makeError(w, InvalidData)
		return
	}
	key := keys[0]

	id, err := strconv.Atoi(key)
	if err != nil {
		makeError(w, InvalidData)
		return
	}

	city := datastore.CityGet(id)
	responseJSON(w, city)
}
func cityGetAllHandler(w http.ResponseWriter, r *http.Request) {
	offsetKeys, ok := r.URL.Query()["offset"]
	if !ok || len(offsetKeys[0]) < 1 {
		makeError(w, InvalidData)
		return
	}
	offsetKey := offsetKeys[0]

	offset, err := strconv.Atoi(offsetKey)
	if err != nil {
		makeError(w, InvalidData)
		return
	}

	limitKeys, ok := r.URL.Query()["limit"]
	if !ok || len(offsetKeys[0]) < 1 {
		makeError(w, InvalidData)
		return
	}
	limitKey := limitKeys[0]

	limit, err := strconv.Atoi(limitKey)
	if err != nil {
		makeError(w, InvalidData)
		return
	}

	var cities []datastore.City
	if limit > 0 {
		cities = datastore.CitiesGetWithOffsetAndLimit(offset, limit)
	} else {
		cities = datastore.CitiesGetAll()
	}

	responseJSON(w, cities)
}
func cityAddHandler(w http.ResponseWriter, r *http.Request) {
	var data datastore.City
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		makeError(w, InvalidData)
		return
	}

	datastore.CityAdd(data)
}
func cityEditHandler(w http.ResponseWriter, r *http.Request) {
	var data datastore.City
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		makeError(w, InvalidData)
		return
	}

	datastore.CityEdit(data)
}
func cityDeleteHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		makeError(w, InvalidData)
		return
	}
	key := keys[0]

	id, err := strconv.Atoi(key)
	if err != nil {
		makeError(w, InvalidData)
		return
	}

	datastore.CityDelete(id)
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin")) // u.Scheme+"://"+u.Host
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-api-token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
