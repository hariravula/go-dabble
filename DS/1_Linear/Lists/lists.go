package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(20)
	intList.PushBack(42)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int)) // This is type assertion. https://tour.golang.org/methods/15
	}
}
