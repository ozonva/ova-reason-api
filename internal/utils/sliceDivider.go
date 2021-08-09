package utils

func DivideToBatches(slice []int, batchSize int) [][]int {
	result := make([][]int, 0, len(slice)/batchSize+1)
	var batch []int
	for index, item := range slice {
		batch = append(batch, item)

		if len(batch) == batchSize || index == len(slice)-1 {
			result = append(result, batch)
			batch = nil
		}
	}

	return result
}
