package main

import (
	"fmt"
)

func visit(numbers []int, callback func(int)) {
	for _, value := range numbers {
		callback(value)
	}
}
func main() {
	i := 0
	fmt.Println("This is awesome")
	fmt.Println("This a tutorial on clousers2!")
	fmt.Println("This a tutorial on clousers!")
	visit([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(n int) {
		fmt.Println(n)
	})
}

