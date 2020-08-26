package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeletePerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	// get the params from the requst
	params := mux.Vars(request)
	// convert params id (string) to MongoDB ID
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("peoplerex").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// get item by id
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	if result.DeletedCount == 0 {
		// log.Fatal("Error on deleting one Hero", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "Unable to delete item"}`))
		return
	}
	response.WriteHeader(http.StatusNoContent)
	json.NewEncoder(response).Encode(result.DeletedCount)
}
