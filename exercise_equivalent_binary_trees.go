package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func Walk(t *tree.Tree, ch chan int) {
	sendValue(t, ch)
	close(ch)
}

func sendValue(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	sendValue(t.Left, ch)
	ch <- t.Value
	sendValue(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
