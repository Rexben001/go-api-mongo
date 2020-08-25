package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

func main() {
	fmt.Println("App has started!!!!")
	// define timeout for Mongo and Go
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// mongodb connection
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if client != nil {
		fmt.Println("Connected successfully")

	}
	// define router
	router := mux.NewRouter()
	router.HandleFunc("/person", AddPerson).Methods("POST")
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
	http.ListenAndServe(":12345", router)
}
