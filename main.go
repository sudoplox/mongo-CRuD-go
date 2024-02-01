package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/sudoplox/mongo-CRuD-go/controllers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"net/http"
	"os"
	"os/signal"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	// creates a new instance
	r := httprouter.New()

	client := getClient()
	// user controller
	uc := controllers.NewUserController(client)

	// SIGINT -> disconnects from Mongo and exits
	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		log.Println("Program killed !")
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		fmt.Println("You successfully disconnected from MongoDB!")
		os.Exit(0)
	}()

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	// creates the golang server
	err := http.ListenAndServe("localhost:9000", r)
	if err != nil {
		fmt.Println(err)
	}
}

func getClient() *mongo.Client {

	// example: mongodb+srv://root:root@cluster-name.qpwoei.mongodb.net/
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_CLUSTER"))
	uri = uri + "?retryWrites=true&w=majority"

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client
}
