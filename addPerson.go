package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func AddPerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var person Person

	// get the body request and decode it
	json.NewDecoder(request.Body).Decode(&person)

	collection := client.Database("peoplerex").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(result)
}
