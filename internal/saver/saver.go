package saver

import (
	"fmt"
	"github.com/ozonva/ova-reason-api/internal/flusher"
	"github.com/ozonva/ova-reason-api/internal/model"
	"strconv"
	"sync"
	"time"
)

type Saver interface {
	Save(entity model.Reason) // заменить на свою сущность
	// Init()
	Close()
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
) Saver {
	newSaver := saver{
		capacity:    capacity,
		flusherObj:  flusher,
		tempStorage: make([]model.Reason, 0, capacity),
	}
	go newSaver.flushByTimeout()
	return &newSaver
}

type saver struct {
	capacity    uint
	flusherObj  flusher.Flusher
	tempStorage []model.Reason
	mu          sync.Mutex
	totalCnt    int
}

func (s *saver) Save(entity model.Reason) {
	s.mu.Lock()
	s.tempStorage = append(s.tempStorage, entity)
	s.mu.Unlock()
}

func (s *saver) Close() {

	s.flush()

	fmt.Println("TotalCnt= " + strconv.Itoa(s.totalCnt))
}

func (s *saver) flush() {

	s.mu.Lock()
	whatToFlush := s.tempStorage
	time.Sleep(time.Millisecond * 500)
	s.tempStorage = make([]model.Reason, 0, s.capacity)
	s.mu.Unlock()

	fmt.Println("Saver: Start flushing " + strconv.Itoa(len(whatToFlush)) + " reasons")
	for ind, reason := range whatToFlush {
		fmt.Println("Saver: " + strconv.Itoa(ind) + " " + reason.Why)
	}
	s.totalCnt += len(whatToFlush)

	if len(whatToFlush) > 0 {
		s.flusherObj.Flush(whatToFlush)
	}
}

func (s *saver) flushByTimeout() {

	for {
		time.Sleep(5 * time.Second)
		go s.flush()
	}

}
