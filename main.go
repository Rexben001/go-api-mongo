package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var client *mongo.Client

func main() {
	fmt.Println("App has started!!!!")

	mongoUri, exists := os.LookupEnv("MONGO_URI")

	if exists {
		fmt.Println("ENV files loaded ")
	}

	// define timeout for Mongo and Go
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// mongodb connection
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	if client != nil {
		fmt.Println("Connected successfully")

	}
	// define router
	router := mux.NewRouter()
	router.HandleFunc("/person", AddPerson).Methods("POST")
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")
	router.HandleFunc("/person/{id}", UpdatePerson).Methods("PUT")
	http.ListenAndServe(":12345", router)
}
