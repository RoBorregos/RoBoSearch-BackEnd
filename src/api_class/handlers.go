package api_class

import (
	"fmt"
	"net/http"
)

func updateCodeHandler(w http.ResponseWriter, r *http.Request, id string) {
	codes, ok := r.URL.Query()["code"]
	if !ok || len(codes[0]) < 1 {
		http.Error(w, "Code param missing.", http.StatusBadRequest)
		return
	}
	// code := codes[0]

	fmt.Fprintf(w, "Soon added %s!", id)
}

func CreateUpdateCodeHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len(path):]
		if len(id) < 1 {
			http.Error(w, "ID missing.", http.StatusBadRequest)
			return
		}
		updateCodeHandler(w, r, id)
	}
}

func addCodeHandler(w http.ResponseWriter, r *http.Request, id string) {
	codes, ok := r.URL.Query()["code"]
	if !ok || len(codes[0]) < 1 {
		http.Error(w, "Code param missing.", http.StatusBadRequest)
		return
	}
	// code := codes[0]

	fmt.Fprintf(w, "Soon added %s!", id)
}

func CreateAddCodeHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len(path):]
		if len(id) < 1 {
			http.Error(w, "ID missing.", http.StatusBadRequest)
			return
		}
		addCodeHandler(w, r, id)
	}
}

func deleteCodeHandler(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprintf(w, "Soon will be gone %s!", id)
}

func CreateDeleteCodeHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len(path):]
		if len(id) < 1 {
			http.Error(w, "ID missing.", http.StatusBadRequest)
			return
		}
		deleteCodeHandler(w, r, id)
	}
}

func getCodeHandler(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprintf(w, "Soon %s!", id)
}

func CreateGetCodeHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len(path):]
		if len(id) < 1 {
			http.Error(w, "ID missing.", http.StatusBadRequest)
			return
		}
		getCodeHandler(w, r, id)
	}
}
