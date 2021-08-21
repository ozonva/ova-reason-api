package repo

import "github.com/ozonva/ova-reason-api/internal/model"

// Repo - интерфейс хранилища для сущности Entity
type Repo interface {
	AddEntities(entities []model.Reason) error
	ListEntities(limit, offset uint64) ([]model.Reason, error)
	DescribeEntity(entityId uint64) (*model.Reason, error)
}
