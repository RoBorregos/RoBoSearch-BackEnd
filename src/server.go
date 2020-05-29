package main

import (
	"fmt"
	"net/http"
	"os"
	"log"

	api_class "./api_class"
)

func main() {
	fmt.Println("Starting")

	var DB string
	if len(os.Getenv("DATABASE_URL")) > 0 {
		DB = os.Getenv("DATABASE_URL")
	} else {
		DB = "mongodb://localhost/"
	}
	if _, err := api_class.InitConnection(DB); err != nil {
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
	http.HandleFunc(
		"/api/class/get-all-code/",
		api_class.CreateGetAllCodeHandler("/api/class/get-all-code/"))

	http.HandleFunc(
		"/api/class/users/check-admin/",
		api_class.CreateCheckIfAdminHandler("/api/class/users/check-admin/"))

	fmt.Println("Started!")
	var PORT string
	if len(os.Getenv("PORT")) > 0 {
		PORT = os.Getenv("PORT")
	} else {
		PORT = "8080"
	}
	log.Fatal(http.ListenAndServe(":" + PORT, nil))
}
