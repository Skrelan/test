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
	i := 'a'
	fmt.Println("This a tutorial on clousers!")
	visit([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(n int) {
		fmt.Println(n)
	})
}

