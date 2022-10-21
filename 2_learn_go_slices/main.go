package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println(fruitList)
	fruitList = append(fruitList, "Banana", "Mango")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:4])
	fmt.Println(fruitList)

	score := make([]int, 4)
	score[0] = 444
	score[1] = 333
	score[2] = 222
	score[3] = 111
	//score[4] = 000
	fmt.Println(score)
	score = append(score, 555, 666, 777, 888)
	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score))
	sort.Ints(score)
	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score))

	// delete a value from slices based on index
	index := 2

	newFruitList := append(fruitList[:index], fruitList[index+1:]...)
	fmt.Println(newFruitList)
}
