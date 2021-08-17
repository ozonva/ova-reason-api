package main

import (
	"fmt"
	"github.com/ozonva/ova-reason-api/internal/utils"
	"os"
	"time"
)

func main() {
	for {
		newConf := readConfig()
		fmt.Printf("Current config: %d\n", newConf)

		time.Sleep(5 * time.Second)
	}
}

func readConfig() int {
	file, err := os.Open("config.txt")
	if err != nil { // если возникла ошибка
		fmt.Println("Unable to open file config.txt: ", err)
		os.Exit(1) // выходим из программы
	}
	defer file.Close()

	var myConf int
	fmt.Fscanf(file, "%d", &myConf)

	return myConf
}

func task2() {
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := utils.DivideToBatches(array[0:], 2)

	fmt.Println(len(result))
	fmt.Println(cap(result))
	for i := range result {
		for j := range result[i] {
			fmt.Printf("result[%d][%d] = %d\n", i, j, result[i][j])
		}
	}
	fmt.Println("___________________________")

	sourceMap := map[int]string{1: "one", 2: "two", 3: "three"}
	result2 := utils.InvertMap(sourceMap)
	for key, value := range result2 {
		fmt.Printf("invertedMap[%s] = %d\n", key, value)
	}

	fmt.Println("___________________________")
	sliceForFilter := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result3 := utils.FilterOdds(sliceForFilter)

	fmt.Printf("Even numbers: ")
	for _, value := range result3 {
		fmt.Printf("%d,", value)
	}
	fmt.Println("\nHello, it's ova-reason-api!")
}
