package component

import (
	"context"
	"fmt"

	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const MongoDBURIString = "mongodb://%s"

func NewMongoDB(config config.MongoDBConfig) (mongodb *mongo.Database, err error) {
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf(MongoDBURIString, config.DBAddress),
	)
	mongo, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return
	}
	mongodb = mongo.Database(
		config.DBName,
		&options.DatabaseOptions{
			ReadPreference: readpref.Nearest(),
		},
	)

	return
}
