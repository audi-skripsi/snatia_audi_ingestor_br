package service

import (
	"sync"

	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/config"
	indto "github.com/audi-skripsi/snatia_audi_ingestor_be/internal/dto"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/model"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/repository"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/pkg/dto"
	"github.com/sirupsen/logrus"
)

type Service interface {
	StoreEvent(event dto.EventLog, microBatch bool) (err error)
}

type service struct {
	logger     *logrus.Entry
	repository repository.Repository
	EventBatch *indto.EventBatch
}

type NewServiceParams struct {
	Logger     *logrus.Entry
	Repository repository.Repository
	Config     *config.Config
}

func NewService(params NewServiceParams) Service {
	s := &service{
		logger:     params.Logger,
		repository: params.Repository,
	}
	s.EventBatch = &indto.EventBatch{
		Mu:             &sync.Mutex{},
		BatchEventData: make([]model.EventData, 0, 101),
	}
	s.initBatchCron()
	return s
}
