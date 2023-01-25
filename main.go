package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gabrielmvas/user-api-golang/http"
	"github.com/gabrielmvas/user-api-golang/repository"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepository(client.Database("users"))

	server := http.NewServer(repository)

	router := gin.Default()
	{
		router.GET("/users/:email", server.GetUser)
		router.GET("/users", server.GetUsers)
		router.POST("/users", server.CreateUser)
		router.PUT("/users/:email", server.UpdateUser)
		router.DELETE("/users/:email", server.DeleteUser)
	}

	router.Run(":9090")
}
