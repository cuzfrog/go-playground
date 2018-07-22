package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func Walk(tree *tree.Tree, c chan int) {
	if tree.Left != nil {
		Walk(tree.Left, c)
	}
	c <- tree.Value
	if tree.Right != nil {
		Walk(tree.Right, c)
	}
}

func HasSameValue(tree1, tree2 *tree.Tree) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)

	go Walk(tree1, c1)
	go Walk(tree2, c2)

	for i := 0; i < 10; i++ {
		if <-c1 != <-c2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(HasSameValue(tree.New(3), tree.New(2)))

	//c := make(chan int, 10)
	//go Walk(tree.New(2), c)
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-c)
	//}
}
