package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("Hello to golang")
	client, ctx, cancel, err := connect("mongodb+srv://hao_admin:hao_admin@sahara.o8tne.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	ping(client, ctx)
}

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("Error!")
		return err
	}
	fmt.Println("Connected Successfully!")
	return nil
}
