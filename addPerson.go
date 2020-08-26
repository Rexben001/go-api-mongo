package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"
)


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
	// writes the objects to standard output
	json.NewEncoder(response).Encode(result)
}
