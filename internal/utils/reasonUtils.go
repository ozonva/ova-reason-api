package utils

import (
	"errors"
	"github.com/ozonva/ova-reason-api/internal/model"
)

func SplitToBulks(reasons []model.Reason, batchSize int) [][]model.Reason {

	result := make([][]model.Reason, 0, len(reasons)/batchSize+1)
	var batch []model.Reason
	for index, item := range reasons {
		batch = append(batch, item)

		if len(batch) == batchSize || index == len(reasons)-1 {
			result = append(result, batch)
			batch = nil
		}
	}

	return result
}

func ConvertToMap(reasons []model.Reason) (map[uint64]model.Reason, error) {
	result := make(map[uint64]model.Reason, len(reasons))

	for _, reason := range reasons {
		if _, found := result[reason.Id]; found {
			return nil, errors.New("duplicated keys")
		}
		result[reason.Id] = reason

	}

	return result, nil
}
