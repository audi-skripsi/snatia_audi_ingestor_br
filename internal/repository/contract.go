package repository

import (
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/config"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/dto"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertEvent(event model.EventData, collName string) (err error)
	MicrobatchInsertEvent(eventBatch *dto.EventBatch) (err error)
}

type repository struct {
	logger *logrus.Entry
	config *repositoryConfig
	mongo  *mongo.Database
}

type repositoryConfig struct {
	mongoConfig config.MongoDBConfig
}

type NewRepositoryParams struct {
	Logger *logrus.Entry
	Config *config.Config
	Mongo  *mongo.Database
}

func NewRepository(params NewRepositoryParams) Repository {
	return &repository{
		logger: params.Logger,
		mongo:  params.Mongo,
		config: &repositoryConfig{
			mongoConfig: params.Config.MongoDBConfig,
		},
	}
}
