package flusher

import (
	"github.com/ozonva/ova-reason-api/internal/model"
	"github.com/ozonva/ova-reason-api/internal/repo"
)

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(entities []model.Reason) []model.Reason
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(chunkSize int, entityRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		repo:      entityRepo,
	}
}

type flusher struct {
	chunkSize int
	repo      repo.Repo
}

func (f flusher) Flush(entities []model.Reason) []model.Reason {
	return entities
}
