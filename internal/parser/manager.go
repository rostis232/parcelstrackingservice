package parser

import (
	"fmt"
	"github.com/rostis232/parcelstrackingservice/models"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type ParsingManager struct {
	mu         sync.Mutex
	timeOut    time.Duration
	queueCount int
	tasksPool  []models.Task
}

func New(timeOut time.Duration, queueCount int) *ParsingManager {
	pm := ParsingManager{
		timeOut:    timeOut,
		queueCount: queueCount,
		tasksPool:  make([]models.Task, 0),
	}

	go pm.checkPoolByTicker()

	return &pm
}

func (pm *ParsingManager) checkPoolByTicker() {
	for range time.Tick(pm.timeOut) {
		pm.mu.Lock()
		if len(pm.tasksPool) == 0 {
			pm.mu.Unlock()
			continue
		}
		tasksPool := pm.tasksPool
		pm.tasksPool = make([]models.Task, 0)
		pm.mu.Unlock()

		go pm.start(tasksPool)
	}
}

func (pm *ParsingManager) AddTask(task models.Task) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.tasksPool = append(pm.tasksPool, task)
	logrus.Infof("New tracking number %s added. Pool size: %d", task.TrackNumber, len(pm.tasksPool))

	if len(pm.tasksPool) >= pm.queueCount {
		tasksPool := pm.tasksPool
		pm.tasksPool = make([]models.Task, 0)
		go pm.start(tasksPool)
	}
}

func (pm *ParsingManager) start(tasksPool []models.Task) {
	inProccessPool := make(map[string][]chan *models.Data)
	trackNumbers := make([]string, 0)

	for _, task := range tasksPool {
		inProccessPool[task.TrackNumber] = append(inProccessPool[task.TrackNumber], task.OutChannel)
		trackNumbers = append(trackNumbers, task.TrackNumber)
	}

	logrus.Infof("starting proccess for numbers: %v", trackNumbers)

	data, err := pm.getInfo(trackNumbers)
	for trackNumber, channels := range inProccessPool {
		result, ok := data[trackNumber]
		if !ok || err != nil {
			for _, channel := range channels {
				channel <- nil
			}
		} else {
			for _, channel := range channels {
				channel <- result
			}
		}
	}
}

func (pm *ParsingManager) getInfo(trackNumbers []string) (map[string]*models.Data, error) {
	const maxRetries = 3

	delay := 500 * time.Millisecond
	for i := 0; i < maxRetries; i++ {
		data, err := pm.tryGetInfo(trackNumbers)
		if err == nil {
			return data, nil
		}

		if i == maxRetries-1 {
			return nil, fmt.Errorf("function failed after %d retries: %w", maxRetries, err)
		}

		time.Sleep(delay)
		delay *= 2
	}

	return nil, nil
}

func (pm *ParsingManager) tryGetInfo(trackNumbers []string) (map[string]*models.Data, error) {
	html, err := ParsePage(trackNumbers)
	if err != nil {
		return nil, err
	}

	data, err := ScrapeData(html)
	if err != nil {
		return nil, err
	}

	return data, nil
}
