package repo

import (
	"fmt"
	"github.com/ozonva/ova-reason-api/internal/model"
)
import "github.com/jmoiron/sqlx"

const tableName = "reasons"

// Repo - интерфейс хранилища для сущности Entity
type Repo interface {
	AddEntity(entities model.Reason) (int64, error)
	ListEntities(limit, offset uint64) ([]model.Reason, error)
	DescribeEntity(entityId uint64) (*model.Reason, error)
	RemoveEntity(entityId uint64) error
	ReplaceEntity(entityId uint64, entities model.Reason) error
}

type ReasonRepository struct {
	db *sqlx.DB
}

func (r ReasonRepository) ReplaceEntity(entityId uint64, entity model.Reason) error {
	_, err := r.db.NamedExec(`UPDATE reasons Set (user_id, action_id, why) = (:userId,:actionId,:why) where id = :id `,
		map[string]interface{}{
			"userId":   entity.UserId,
			"actionId": entity.ActionId,
			"why":      entity.Why,
			"id":       entityId,
		})

	return err
}

func (r ReasonRepository) AddEntity(entity model.Reason) (int64, error) {
	_, err := r.db.NamedExec(`INSERT INTO reasons (user_id, action_id, why) VALUES (:userId,:actionId,:why)`,
		map[string]interface{}{
			"userId":   entity.UserId,
			"actionId": entity.ActionId,
			"why":      entity.Why,
		})
	if err != nil {
		return -1, err
	}

	//commented because "LastInsertId is not supported by this driver"
	//return res.LastInsertId()

	return 0, err
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
	var reason model.Reason
	err := r.db.Get(&reason, "select * from reasons where id = $1 ", entityId)
	if err != nil {
		return nil, err
	}

	return &reason, nil
}

func (r ReasonRepository) RemoveEntity(entityId uint64) error {
	_, err := r.db.NamedExec(`delete from reasons where id = :id`,
		map[string]interface{}{
			"id": entityId,
		})
	return err
}

func NewReasonRepository(db *sqlx.DB) Repo {
	return &ReasonRepository{
		db: db,
	}
}
