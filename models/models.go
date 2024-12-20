package models

import "time"

type AppConfig struct {
	Port          string
	QueueTimeOut  time.Duration
	MaxQueueCount int
}

type Data struct {
	OriginCountry      string       `json:"origin_country"`
	DestinationCountry string       `json:"destination_country"`
	Checkpoints        []Checkpoint `json:"checkpoints"`
}

type Checkpoint struct {
	Status string `json:"status"`
	Date   string `json:"date"`
}

type Task struct {
	TrackNumber string
	OutChannel  chan *Data
}
