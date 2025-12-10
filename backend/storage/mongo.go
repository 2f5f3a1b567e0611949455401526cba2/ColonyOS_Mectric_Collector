// Package storage ColonyOS_Metric_Collector/backend/storage/mongo.go
package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MetricsCollection *mongo.Collection

func InitMongo(uri, dbName, collectionName string) error {
	ctx, cancel := context.WithTimeout(context.Background(),
		10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf(" failed connection to client: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf(" failed pinging: %w", err)
	}
	log.Println(" mongo connection success")

	MongoClient = client
	MetricsCollection = client.Database(dbName).Collection(collectionName)

	return nil
}
