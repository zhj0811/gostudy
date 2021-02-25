package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Buddhist struct {
	Id     primitive.ObjectID `bson:"_id"`
	Name   string             `json:"name"`
	Remark string             `json:"remark"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	buddhist := client.Database("testdb").Collection("buddhist")
	_, err = buddhist.InsertOne(context.Background(), &Buddhist{Id: primitive.NewObjectID(), Name: "TESt", Remark: "test mark"})
	if err != nil {
		log.Fatal(err)
	}
	//context.TODO()
	fmt.Println("Insert to MongoDB!")

	var res Buddhist
	err = buddhist.FindOne(context.TODO(), bson.M{"name": "TESt"}).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("find one res: %v\n", res)
	fmt.Println("find one res id", res.Id.Hex())
	//var res2 Buddhist
	////err = buddhist.FindOne(context.TODO(), bson.M{"_id": res.Id}).Decode(&res2)
	//err = buddhist.FindOne(context.TODO(), &res).Decode(&res2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("find one res2: %v", res2)
	cur, err := buddhist.Find(context.TODO(), bson.M{"name": "TEST"})
	if err != nil {
		log.Fatal(err)
	}
	var res2 []Buddhist
	err = cur.All(context.TODO(), &res2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("find res: %v", res2)
	count, err := buddhist.CountDocuments(context.TODO(), bson.M{"name": "TEST"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("find res count: %d", count)
}
