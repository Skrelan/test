package main

import (
	"fmt"
)

func visit(numbers []int, callback func(int)) {
	for _, value := range numbers {
		callback(value)
	}
}
func branched(){
	fmt.Println("this is part of a feature branch")
}
func uptow(){
	panic()
}
func main() {
	i := 'a'
	fmt.Println("This is awesome",i)
	fmt.Println("This a tutorial on clousers2! on Branched")
	fmt.Println("This a tutorial on clousers!")
	visit([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(n int) {
		fmt.Println(n)
	})
}

