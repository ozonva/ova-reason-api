package oldTasks

import (
	"errors"
	"fmt"
	"github.com/ozonva/ova-reason-api/internal/model"
	"github.com/ozonva/ova-reason-api/internal/utils"
	"os"
	"time"
)

func task3() {

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
		newConf, err := readConfig()
		if err == nil {
			fmt.Printf("Current config: %d\n", newConf)
		} else {
			fmt.Println(err.Error())
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
