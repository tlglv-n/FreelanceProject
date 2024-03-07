package market

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// mongodb://username:password@localhost:27017/?retryWrites=true&w=majority&tls=false

const timeout = 10 * time.Second

type Mongo struct {
	Client *mongo.Client
}

func NewMongo(url string) (market Mongo, err error) {
	market.Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err = market.Client.Connect(ctx); err != nil {
		return
	}

	if err = market.Client.Ping(context.Background(), nil); err != nil {
		return
	}

	return
}
