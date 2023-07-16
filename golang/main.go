package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(context)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context)

	collection := client.Database("cooker").Collection("recipes")
	filters := bson.M{"tags": bson.M{"$in": []string{"easy", "fish"}}}
	opts := options.Find()
	opts.SetSort(bson.M{"title": 1})
	opts.SetLimit(3)
	cur, err := collection.Find(context, filters, opts)
	defer cur.Close(context)
	if err != nil {
		log.Fatal(err)
	}

	var recipes []struct {
		Title string   `bson:"title,omitempty"`
		Tags  []string `bson:"tags,omitempty"`
	}
	err = cur.All(context, &recipes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(recipes)
}
