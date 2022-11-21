package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Akbar")
	data.PushBack("Babelan")
	data.PushBack("Musthofa")
	data.PushFront("Kholifah")

	//data.Front().Next().Next()

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	//for element := data.Back(); element != nil; element.Prev() {
		//fmt.Println(element.Value)
	}
}
