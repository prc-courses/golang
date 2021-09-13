package main

import "fmt"

import "container/list"

func insertListElements(n int) *list.List {
	lst := list.New()
	for i := 1; i <= n; i++ {
		lst.PushBack(i)
	}
	return lst
}

func main() {
	n := 5
	myList := insertListElements(n)
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
