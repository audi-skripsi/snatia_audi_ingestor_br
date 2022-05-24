package dto

type EventLog struct {
	Level     string `json:"level"`
	AppName   string `json:"app_name"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
