package saver

import (
	"github.com/ozonva/ova-reason-api/internal/flusher"
	"github.com/ozonva/ova-reason-api/internal/model"
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
	return &saver{
		capacity: capacity,
		flusher:  flusher,
	}
}

type saver struct {
	capacity uint
	flusher  flusher.Flusher
}

func (s saver) Save(entity model.Reason) {
	panic("implement me")
}

func (s saver) Close() {
	panic("implement me")
}
