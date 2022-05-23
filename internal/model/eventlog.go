package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventData struct {
	ObjectID  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Level     string             `json:"level" bson:"level"`
	AppName   string             `json:"app_name" bson:"app_name"`
	Message   string             `json:"Message" bson:"data"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}
