package main

import (
	"fmt"

	"github.com/ozonva/ova-reason-api/internal/utils"
)

func main() {

	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := utils.DivideToBatches(array[0:], 2)

	fmt.Println(len(result))
	fmt.Println(cap(result))
	for i := range result {
		for j := range result[i] {
			fmt.Printf("result[%d][%d] = %d\n", i, j, result[i][j])
		}
	}

	fmt.Println("Hello, it's ova-reason-api!")
}
