package service

import (
	"time"

	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/util/converterutil"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/pkg/dto"
)

func (s *service) StoreEvent(event dto.EventLog, microBatch bool) (err error) {
	if microBatch {
		err = s.processMicrobatch(event)
	} else {
		err = s.processNonMicrobatch(event)
	}

	if err != nil {
		s.logger.Errorf("error processing data of %+v: %+v", event, err)
	}

	return
}

func (s *service) processNonMicrobatch(event dto.EventLog) (err error) {
	eventModel := converterutil.EventDtoToModel(event)
	err = s.repository.InsertEvent(eventModel, "non_microbatch")
	return
}

func (s *service) processMicrobatch(event dto.EventLog) (err error) {
	eventModel := converterutil.EventDtoToModel(event)
	s.EventBatch.BatchEventData = append(s.EventBatch.BatchEventData, eventModel)

	if len(s.EventBatch.BatchEventData) == 500 {
		s.EventBatch.Mu.Lock()
		err = s.repository.MicrobatchInsertEvent(s.EventBatch, "microbatch")
		s.EventBatch.BatchEventData = nil
		s.EventBatch.Mu.Unlock()
	}

	return
}

func (s *service) initBatchCron() {
	go func() {
		for {
			if len(s.EventBatch.BatchEventData) > 0 {
				s.EventBatch.Mu.Lock()
				err := s.repository.MicrobatchInsertEvent(s.EventBatch, "microbatch")
				if err != nil {
					s.logger.Errorf("error batch insert: %+v", err)
				}
				s.EventBatch.BatchEventData = nil
				s.EventBatch.Mu.Unlock()
			}
			time.Sleep(2 * time.Second)
		}
	}()
}
