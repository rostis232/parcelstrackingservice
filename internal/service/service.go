package service

import (
	"errors"
	"github.com/rostis232/parcelstrackingservice/internal/parser"
	"github.com/rostis232/parcelstrackingservice/models"
	"time"
)

type ParsingManager interface {
	AddTask(task models.Task)
}

type Service struct {
	parsingManager ParsingManager
}

func New(timeOut time.Duration, maxQueueCount int) *Service {
	return &Service{
		//TODO
		parsingManager: parser.New(timeOut, maxQueueCount),
	}
}

func (s *Service) GetInfo(trackNumber string) (*models.Data, error) {
	resultChannel := make(chan *models.Data)
	defer close(resultChannel)

	task := models.Task{
		TrackNumber: trackNumber,
		OutChannel:  resultChannel,
	}

	s.parsingManager.AddTask(task)

	result := <-resultChannel
	if result == nil {
		return nil, errors.New("something went wrong. Check trackNumber")
	}
	return result, nil
}
