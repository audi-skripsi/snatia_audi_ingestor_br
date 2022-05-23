package dto

import (
	"sync"

	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/model"
)

type EventBatch struct {
	Mu *sync.Mutex

	CollName       string
	BatchEventData []model.EventData
}
