package main

import (
	"fmt"
	"net/http"
	"log"

	api_class "./api_class"
)

func main() {
	fmt.Println("Starting")
	if _, err := api_class.InitConnection("mongodb://localhost/"); err != nil {
		log.Fatal("Connection db failed: %s", err.Error())
		return
	}

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

	fmt.Println("Started!")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
