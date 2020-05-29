package api_class

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func updateCodeHandler(w http.ResponseWriter, r *http.Request, id string) {
	var elms bson.D
	fds := 0

	codes, ok := r.URL.Query()["code"]
	if ok && len(codes[0]) > 0  {
		elms = append(elms, bson.E{"code", codes[0]})
		fds += 1
	}
	filenames, ok := r.URL.Query()["filename"]
	if ok && len(filenames[0]) > 0 {
		elms = append(elms, bson.E{"filename", filenames[0]})
		fds += 1
	}
	if fds == 0 {
		http.Error(w, "At leats code or filename should be sent.",
		 http.StatusBadRequest)
		return
	}
	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID doesn't exist.", http.StatusNotFound)
		return
	}

	filter := bson.M{"_id": obj_id}
	update := bson.D{{"$set", elms}}
	result, err := Database.Collection("codes").UpdateOne(
		context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if result.MatchedCount == 0 {
		http.Error(w, "ID doesn't exist.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Updated!")
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

func addCodeHandler(w http.ResponseWriter, r *http.Request) {
	codes, ok := r.URL.Query()["code"]
	if !ok || len(codes[0]) < 1 {
		http.Error(w, "Code param missing.", http.StatusBadRequest)
		return
	}
	code := codes[0]
	filenames, ok := r.URL.Query()["filename"]
	if !ok || len(filenames[0]) < 1 {
		http.Error(w, "Filename param missing.", http.StatusBadRequest)
		return
	}
	filename := filenames[0]

	res, err := Database.Collection("codes").InsertOne(
		context.Background(), bson.M{"code": code, "filename": filename})
	if err != nil { 
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if obj_id, ok := res.InsertedID.(primitive.ObjectID); ok { 
		jres := &Code{Id: obj_id, Filename: filename, Code: code}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jres)
	} else {
		http.Error(w, "Error converting an ID.", http.StatusInternalServerError)
	}
}

func CreateAddCodeHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addCodeHandler(w, r)
	}
}

func deleteCodeHandler(w http.ResponseWriter, r *http.Request, id string) {
	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID doesn't exist.", http.StatusNotFound)
		return
	}

	res, err := Database.Collection("codes").DeleteOne(
		context.Background(), bson.M{"_id": obj_id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if res.DeletedCount == 0 {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Deleted it!")
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
	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID doesn't exist.", http.StatusNotFound)
		return
	}
	filter := bson.M{"_id": obj_id}

	result := &Code{}
	err = Database.Collection("codes").FindOne(
		context.Background(), filter).Decode(result)

	if err == mongo.ErrNoDocuments {
		http.Error(w, "ID doesn't exist.", http.StatusNotFound)
		return
	}
	if err != nil { 
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
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
