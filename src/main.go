package main

import (
	"DS"
)

func main() {
	list := new(DS.Linklist)
	list.Init([]int{1, 2, 3, 4, 6, 7})
	list.Show()
	list.Reverse()
	list.Show()
}
