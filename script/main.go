package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDataDto struct {
	token       string
	carrierInfo interface{}
	updateTime  string
}

func main() {
	client1 := DB1()
	client2 := DB2()
	//查询第一个
	collection := client1.Database("mh-data").Collection("mhData")
	collection2 := client2.Database("mh-data").Collection("mhData")
	//插入
	for {
		var results bson.M
		collection.FindOneAndDelete(context.TODO(), bson.M{"updateTime": bson.M{"$gte": "2021-04-20 10:38:25"}}).Decode(&results)
		fmt.Println("token:", results["token"])
		fmt.Println("updateTime:", results["updateTime"])

		one, err := collection2.InsertOne(context.TODO(), results)

		fmt.Println(one, err)
		if err != nil {
			break
		}
	}
}

func DB1() *mongo.Client {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

func DB2() *mongo.Client {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27018")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
