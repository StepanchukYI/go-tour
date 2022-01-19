package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chan1, chan2 := make(chan int), make(chan int)
	go Walk(t1, chan1)
	go Walk(t2, chan2)

	for i := range chan1 {
		fmt.Println(i)
	}
	for j := range chan2 {
		fmt.Println(j)
	}

	for {
		num1, ok1 := <-chan1
		num2, ok2 := <-chan2
		fmt.Println(num1, ok1)
		fmt.Println(num2, ok2)
		if ok1 != ok2 || num1 != num2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(t1)
	fmt.Println(t2)
	go Walk(t1, ch)
	fmt.Println("Walked")
	fmt.Println(ch)
	fmt.Println(Same(t1, t1))
	fmt.Println(Same(t1, t2))

}
