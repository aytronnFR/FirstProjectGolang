package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Translation struct {
	ID           primitive.ObjectID `bson:"_id"`
	Language     string             `bson:"language"`
	Key          string             `bson:"key"`
	Translations string             `bson:"translations"`
}

func getMongo(collectionName string) {

	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("Zakary").Collection(collectionName)

	var tasks []*Translation

	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		info("ERROR FIND")
		return
	}

	for cur.Next(ctx) {
		var t Translation
		err := cur.Decode(&t)
		if err != nil {
			fmt.Println(tasks)
			return
		}

		tasks = append(tasks, &t)
	}

	if err := cur.Err(); err != nil {
		fmt.Println(tasks)
		return
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(tasks) == 0 {
		fmt.Println(tasks)
		return
	}

	for _, entry := range tasks {
		fmt.Println(entry.Key)
	}
}
