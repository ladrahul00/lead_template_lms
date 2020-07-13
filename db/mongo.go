package db

import (
	"context"

	"github.com/wolf00/lead_template_lms/constants"

	log "github.com/micro/go-micro/v2/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongo(ctx context.Context) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(constants.MongoConnectionString))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to MongoDB!...")

	return client.Database(constants.DatabaseName)
}

// LeadTemplates used to retrieve collection from database
func LeadTemplates(ctx context.Context) *mongo.Collection {
	return connectMongo(ctx).Collection(constants.LeadTemplates)
}
