package main

import(
	"container/list"
	"fmt"
)

func main() {
	mainList()
}

func mainList() {
	l := list.New()
	l.PushBack("B")
	l.PushBack("C")
	l.PushBack("D")
	l.PushFront("A")

	fmt.Println("Length:", l.Len())

	//iterate forward
	fmt.Println("Iteration Forward")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	
	//iterate backward
	fmt.Println("Iteration Backward")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	k := list.New()
	k.PushBack(1)
	k.PushBack(2)

	// insert before
	a := k.PushBack(4)
	k.InsertBefore(3, a)
	
	// insert after
	b := k.PushBack(5)
	k.InsertAfter(6, b)

	for e := k.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// move item
	fmt.Println("move before function")
	first := k.PushFront(0)
	k.MoveBefore(first, k.Front())
	for e := k.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("move after function")
	last := k.PushBack(7)
	k.MoveAfter(last, k.Back())
	for e := k.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("remove function")
	k.Remove(k.Front())

	for e := k.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}