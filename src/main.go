package main

import (
	"DS"
	"fmt"
)

func main() {
	list := new(DS.Linklist)
	list.Init([]int{1, 2, 3, 4, 6, 7})
	list.Show()
	list.Reverse()
	list.Show()
	list.Insert(7, 123)
	list.Show()
	c := list.Delete(2)
	fmt.Println(c)
	list.Show()
	list.Push(999)
	list.Show()
	fmt.Println(list.GetLength())
	list.Pop()
	list.Show()
}
