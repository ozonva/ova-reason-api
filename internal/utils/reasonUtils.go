package utils

import (
	"errors"
	"github.com/ozonva/ova-reason-api/internal/model"
)

func SplitToBulks(reasons []model.Reason, batchSize int) [][]model.Reason {

	resultLen := calcBulksCnt(len(reasons), batchSize)
	result := make([][]model.Reason, 0, resultLen)

	for i := 0; i < resultLen; i++ {
		startInd := i * batchSize
		endInd := (i + 1) * batchSize
		if endInd > len(reasons) {
			endInd = len(reasons)
		}
		result = append(result, reasons[startInd:endInd])
	}

	return result
}

func calcBulksCnt(arrayLen int, batchSize int) int {
	result := 0
	if arrayLen%batchSize == 0 {
		result = arrayLen / batchSize
	} else {
		result = arrayLen/batchSize + 1
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
