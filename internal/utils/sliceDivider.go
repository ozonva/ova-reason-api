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

func InvertMap(sourceMap map[int]string) map[string]int {
	result := make(map[string]int, len(sourceMap))
	for key, value := range sourceMap {
		result[value] = key
	}

	return result
}

var oddNumbers = []int{1, 3, 5, 7, 9, 11, 13, 15}

func FilterOdds(slice []int) []int {
	var result []int
	for _, value := range slice {

		if !findElement(oddNumbers, value) {
			result = append(result, value)
		}
	}

	return result
}

func findElement(slice []int, element int) bool {
	for _, elem := range slice {
		if element == elem {
			return true
		}

	}
	return false
}
