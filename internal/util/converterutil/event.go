package converterutil

import (
	"time"

	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/model"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/pkg/dto"
)

func EventDtoToModel(in dto.EventLog) model.EventData {
	return model.EventData{
		Level:     in.Level,
		AppName:   in.AppName,
		Message:   in.Message,
		Timestamp: time.Unix(in.Timestamp, 0),
	}
}
