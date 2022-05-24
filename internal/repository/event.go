package repository

import (
	"context"

	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/dto"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/model"
)

func (r *repository) InsertEvent(event model.EventData, collName string) (err error) {
	coll := r.mongo.Collection(collName)

	_, err = coll.InsertOne(context.Background(), event)

	if err != nil {
		r.logger.Errorf("error inserting to mongodb for %+v: %+v", event, err)
	}
	return
}

func (r *repository) MicrobatchInsertEvent(eventBatch *dto.EventBatch, collName string) (err error) {
	coll := r.mongo.Collection(collName)
	var documents []interface{}
	for _, v := range eventBatch.BatchEventData {
		documents = append(documents, v)
	}

	_, err = coll.InsertMany(context.Background(), documents)
	if err != nil {
		r.logger.Errorf("error inserting to mongodb for %+v: %+v", documents, err)
	}

	return
}
