package main

import (
	"github.com/audi-skripsi/snatia_audi_ingestor_be/cmd/http"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/component"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/config"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/repository"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/service"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/pkg/util/logutil"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	config.Init()

	conf := config.Get()

	var level logrus.Level
	if conf.Environment == config.EnvDev {
		level = logrus.DebugLevel
	} else if conf.Environment == config.EnvProd {
		level = logrus.WarnLevel
	}

	logger := logutil.NewLogger(logutil.NewLoggerParams{
		PrettyPrint: true,
		ServiceName: conf.AppName,
		Level:       level,
	})

	logger.Infof("app initialized with config of :%+v", conf)

	ec := echo.New()
	ec.HideBanner = true

	mongo, err := component.NewMongoDB(conf.MongoDBConfig)
	if err != nil {
		logger.Fatalf("[main] error initializing mongodb: %+v", err)
	}

	repository := repository.NewRepository(repository.NewRepositoryParams{
		Logger: logger,
		Config: conf,
		Mongo:  mongo,
	})

	service := service.NewService(service.NewServiceParams{
		Logger:     logger,
		Repository: repository,
		Config:     conf,
	})

	http.StartServer(&http.ServerInitParams{
		Logger:  logger,
		Config:  conf,
		Ec:      ec,
		Service: service,
	})
}
