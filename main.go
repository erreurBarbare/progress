package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Learning struct {
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
}

func (l *Learning) setLearning(category, description string) {
	l.Category = category
	l.Description = description
	l.Created = time.Now()
}

func (l *Learning) setLearningWithTime(category, description string, created time.Time) {
	l.Category = category
	l.Description = description
	l.Created = created
}

func (l *Learning) setCategory(category string) {
	l.Category = category
}

func (l *Learning) setDescription(description string) {
	l.Description = description
}

func main() {
	connectDB("mongodb://localhost:27017/")

	l1 := Learning{}
	l1.setLearning("garden", "plant tomatoes")
	res1A, _ := json.Marshal(l1)
	fmt.Println(string(res1A))

}

func connectDB(url string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	mongoClient, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(url),
	)

	defer func() {
		cancel()
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("mongodb disconnect error : %v", err)
		}
	}()

	if err != nil {
		log.Fatalf("connection error :%v", err)
		return
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping mongodb error :%v", err)
		return
	}
	fmt.Println("ping success")
}
