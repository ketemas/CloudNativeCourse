package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongodbEndpoint = "webserver://172.17.0.2:27017" // Find this from the Mongo container
)

func(d dollars) String() string {return fmt.Sprintf("$#{d} \n")}

type item struct {
	ID			primitive.ObjectID `bson:"_id"`
	Product		string             `bson:"product"`
	Price		string             `bson:"price"`

}

func list(w http.ResponseWriter, r *http.Request){

	db := db.DB().Database(name: "myDB").collection(name: "products")
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
			fmt.Fprintf(w, "Item: #{items, Product}\tprice: #{items.Price}\n")
		}
	}

}
 
func price(w http.ResponseWriter, r *http.Request){
	product := r.URL.Query().Get(key: "product")
	items := Items{}
	db := db.DB().Database(name: "myDB").collection(name: "products")
	err := db.FindOne(context.TODO(), bson.M{"product": bson.M{"$eq": product}}).Decode(&items)
	if err != nil {
		w.WriteHeader(http.statusNotFound)
		fmt.Fprintf(w, format: "error Listing Items.")
	}

	if err != nil {
		w.WriteHeader(http.statusNotFound)
		fmt.Fprintf(w, format: "error Listing Item decode")
	}else {
		fmt.Fprintf(w, "item: #{items.Product}\tPrice: #{items.Price}\n")
	}
}

func update(w http.ResponseWriter, r *http.Request){
	product := r.URL.Query().Get(key: "product")
	price := r.URL.Query().Get(key: "price")
	db := db.DB().Database(name: "myDB").collection(name: "products")
	filter := bson.M{"product" :bson.M{"seq": product}} 	//converting 
	after := options.after
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDucuments: &after,
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "price", Value: price}}}}
	updateResult := db.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)
	fmt.Println(updateResult)
}

func create(w http.ResponseWriter, r *http.Request){
	product := r.URL.Query().Get(key: "product")
	price := r.URL.Query().Get(key: "price")
	db := db.DB().Database(name: "myDB").collection(name: "products")
	res, _ := db.InsertOne(context.TODO(), &Item{
		ID:			primitive.NewObjectID(),
		Product:	product,
		Price:		price,
	})
	fmt.Fprintf(w, "#{items.Product}\tPrice: #{items.Price}\t added to database\n")
}
