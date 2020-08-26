package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	database, _ := os.LookupEnv("DATABASE_NAME")

	var person Person
	// get the params from the requst
	params := mux.Vars(request)
	// convert params id (string) to MongoDB ID
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database(database).Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// get item by id
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(person)
}
