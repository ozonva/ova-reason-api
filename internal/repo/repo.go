package repo

import (
	"fmt"
	"github.com/ozonva/ova-reason-api/internal/model"
)
import "github.com/jmoiron/sqlx"

const tableName = "reasons"

// Repo - интерфейс хранилища для сущности Entity
type Repo interface {
	AddEntities(entities []model.Reason) error
	ListEntities(limit, offset uint64) ([]model.Reason, error)
	DescribeEntity(entityId uint64) (*model.Reason, error)
	RemoveEntity(entityId uint64) error
}

type ReasonRepository struct {
	db *sqlx.DB
}

func (r ReasonRepository) AddEntities(entities []model.Reason) error {
	panic("implement me")
}

func (r *ReasonRepository) ListEntities(limit, offset uint64) ([]model.Reason, error) {
	sql := fmt.Sprintf("select * from reasons LIMIT %d OFFSET %d ", limit, offset)
	var reasons []model.Reason
	err := r.db.Select(&reasons, sql)
	if err != nil {
		return nil, err
	}

	return reasons, nil
}

func (r ReasonRepository) DescribeEntity(entityId uint64) (*model.Reason, error) {
	panic("implement me")
}

func (r ReasonRepository) RemoveEntity(entityId uint64) error {
	panic("implement me")
}

func NewReasonRepository(db *sqlx.DB) Repo {
	return &ReasonRepository{
		db: db,
	}
}
