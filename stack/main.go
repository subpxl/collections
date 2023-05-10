package main

import "fmt"

// lifo
func main() {

	var mys Stack
	fmt.Println(mys.isEmpty())
	mys.Push(99)
	mys.Push(1)
	mys.Push(2)
	mys.Push(3)
	mys.Push(4)
	fmt.Println(mys.isEmpty())
	fmt.Println(mys.Pop())
	fmt.Println(mys.Pop())
	fmt.Println(mys.Peek())

}

type Stack struct {
	Items []int
}

func (s *Stack) isEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack) Push(data int) {
	s.Items = append(s.Items, data)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.Items) == 0 {
		return 0, false
	} else {
		lastItem := len(s.Items) - 1
		s.Items = s.Items[:len(s.Items)-1]
		return lastItem, true
	}
}

func (s *Stack) Peek() int {
	if len(s.Items) == 0 {
		fmt.Println("no stack")
		return 0
	} else {
		lastItem := len(s.Items) - 1
		return lastItem
	}
}
