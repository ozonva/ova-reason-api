package main

import (
	"errors"
	"fmt"
	"github.com/ozonva/ova-reason-api/internal/model"
	"github.com/ozonva/ova-reason-api/internal/utils"
	"os"
	"time"
)

func main() {

	//Task C
	task3C()

	//Task B
	task3B()

	//Task A
	task3A()
}

func readConfig() (int, error) {
	file, err := os.Open("config.txt")
	if err != nil { // если возникла ошибка
		return 0, errors.New("Unable to open file config.txt: " + err.Error())
	}
	defer file.Close()

	var myConf int
	fmt.Fscanf(file, "%d", &myConf)

	return myConf, nil
}

func task3A() {
	fmt.Println("Task 3A")
	for {
		newConf, error := readConfig()
		if error == nil {
			fmt.Printf("Current config: %d\n", newConf)
		} else {
			fmt.Println(error.Error())
		}

		time.Sleep(5 * time.Second)
	}
}

func task3C() {
	fmt.Println("Task 3C")

	slice := []model.Reason{
		*model.New(5, 5, 5, "why5"),
		*model.New(6, 6, 6, "why6"),
		*model.New(7, 7, 7, "why7"),
	}

	reasonMap, err := utils.ConvertToMap(slice)
	if err != nil {
		fmt.Println(err)
	}

	for _, reason := range reasonMap {
		fmt.Println(reason.String())
	}

	bulks := utils.SplitToBulks(slice, 2)
	for i := range bulks {
		for j := range bulks[i] {
			fmt.Printf("result[%d][%d] = %s\n", i, j, bulks[i][j].String())
		}
	}

}

func task3B() {
	fmt.Println("Task 3B")
	s := model.Reason{
		Id:       1,
		UserId:   2,
		ActionId: 3,
		Why:      "qqq",
	}

	fmt.Println(s.String())
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
