package api_class

import (
	"context"
	"encoding/json"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func checkIfAdminHandler(w http.ResponseWriter, r *http.Request, id string) {
	type Result struct{Admin bool `json:"admin"`};

	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&Result{Admin: false})
		return
	}

	filter := bson.M{"_id": obj_id, "admin": 1}
	var dummy struct{}
	err = Database.Collection("admins").FindOne(
		context.Background(), filter).Decode(&dummy)

	if err == mongo.ErrNoDocuments {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&Result{Admin: false})
		return
	}
	if err != nil { 
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Result{Admin: true})
}

func CreateCheckIfAdminHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*");
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization");
    w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE");

		id := r.URL.Path[len(path):]
		if len(id) < 1 {
			http.Error(w, "ID missing.", http.StatusBadRequest)
			return
		}
		checkIfAdminHandler(w, r, id)
	}
}
