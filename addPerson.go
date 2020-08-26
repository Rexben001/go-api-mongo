package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InsertOneResult struct {
	// The identifier that was inserted.
	InsertedID interface{}
}
type ObjectID [12]byte

func AddPerson(response http.ResponseWriter, request *http.Request) {

	database, _ := os.LookupEnv("DATABASE_NAME")

	response.Header().Add("content-type", "application/json")
	var person Person

	// get the body request and decode it
	//json.NewDecoder() removes all but the Name field from each object
	json.NewDecoder(request.Body).Decode(&person)

	collection := client.Database(database).Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	var id string

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		id = oid.Hex()
	}
	finalResult := make(map[string]string)

	finalResult["message"] = "New person added successfully"
	finalResult["InsertedId"] = id
	finalResult["status"] = "200"
	finalResult["success"] = "True"

	// writes the objects to standard output
	json.NewEncoder(response).Encode(finalResult)
}
