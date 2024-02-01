package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/sudoplox/mongo-CRuD-go/models"
	mgoBson "gopkg.in/mgo.v2/bson"
	"os"

	mdBson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func NewUserController(c *mongo.Client) *UserController {
	return &UserController{
		Client: c,
	}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !mgoBson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := mgoBson.ObjectIdHex(id)
	u := models.User{}

	if err := uc.Client.
		Database(os.Getenv("DB_NAME")).
		Collection(os.Getenv("DB_COLLECTION")).
		FindOne(context.TODO(), mdBson.M{"_id": oid}).
		Decode(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set(
		"Content-Type", "application/json",
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
		return
	}
	u.Id = mgoBson.NewObjectId()

	insertOneResult, err := uc.Client.
		Database(os.Getenv("DB_NAME")).
		Collection(os.Getenv("DB_COLLECTION")).
		InsertOne(context.TODO(), u)
	if err != nil {
		fmt.Println(err)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set(
		"Content-Type", "application/json",
	)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Inserted Id: %s\nInserted Obj: %s\n", insertOneResult.InsertedID, uj)

}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !mgoBson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := mgoBson.ObjectIdHex(id)
	u := models.User{}

	if err := uc.Client.
		Database(os.Getenv("DB_NAME")).
		Collection(os.Getenv("DB_COLLECTION")).
		FindOneAndDelete(context.TODO(), mdBson.M{"_id": oid}).
		Decode(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted user: %s\n", uj)
}
