package model

import (
	"fmt"
)

type Reason struct {
	Id       uint64 `db:"id"`
	UserId   uint64 `db:"user_id"`
	ActionId uint64 `db:"action_id"`
	Why      string `db:"why"`
}

func New(userId uint64, reasonId uint64, actionId uint64, why string) *Reason {
	return &Reason{
		UserId:   userId,
		Id:       reasonId,
		ActionId: actionId,
		Why:      why,
	}
}

func (reason *Reason) String() string {
	return fmt.Sprintf("Reason: id= %v, userId=%v, acionId= %v, why= %s ", reason.Id, reason.UserId, reason.ActionId, reason.Why)
}
