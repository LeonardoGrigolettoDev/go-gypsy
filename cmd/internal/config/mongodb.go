package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// ConnectMongoDB configura e conecta ao MongoDB
func ConnectMongoDB(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Verifica a conexão
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	MongoClient = client
	log.Println("Conexão com MongoDB estabelecida com sucesso!")
	return nil
}

// GetCollection retorna uma coleção do MongoDB
func GetCollection(database, collection string) *mongo.Collection {
	return MongoClient.Database(database).Collection(collection)
}
