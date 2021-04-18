//db.recipedb.insert ({ name: " ", country: " ", ingredient: [" "], instruction: [" "] })

package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	recipeEndpoint = "recipedb://172.17.0.5:27017" // Find this from the Mongo container
)

type item struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Country      string             `bson:"country"`
	Recipe      []string           `bson:"recipe"`
	Ingredient       []string           `bson:"ingredient "`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func list(w http.ResponseWriter, r *http.Request){

	db := db.DB().Database(name: "recipeDB").collection(name: "recipes")
	foundItems, err := db.Find(context.TODO(), bson.D{})
	if err != nil {
		w.WriteHeader(http.statusNotFound)
		fmt.Fprintf(w, format: "error Listing Items.")
	}
	for foundItems.Next(context.TODO()){
		items := Items{}
		err := foundItems.Decode(&items)
		if err != nil {
			w.WriteHeader(http.statusNotFound)
			fmt.Fprintf(w, format: "error Listing Item decode")
		}else {
			fmt.Fprintf(w, "Item: #{items, Name}\tIngredient: #{items.Ingredient}\tRecipe: #{items.Recipe}\n")
		}
	}
}

