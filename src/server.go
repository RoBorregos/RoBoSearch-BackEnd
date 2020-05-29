package main

import (
	"net/http"
	"log"
	api_class "./api_class"
)

func main() {
	http.HandleFunc(
		"/api/class/get-code/",
		api_class.CreateGetCodeHandler("/api/class/get-code/"))
	http.HandleFunc(
		"/api/class/update-code/",
		api_class.CreateUpdateCodeHandler("/api/class/update-code/"))
	http.HandleFunc(
		"/api/class/delete-code/",
		api_class.CreateDeleteCodeHandler("/api/class/delete-code/"))
	http.HandleFunc(
		"/api/class/add-code/",
		api_class.CreateAddCodeHandler("/api/class/add-code/"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
