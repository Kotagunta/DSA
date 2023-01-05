package main

import "fmt"

func LinearSearch(items []int, num int) {
	for index, item := range items {
		if item == num {
			fmt.Println("Item found at position =", index+1)
			return
		}
	}
	fmt.Println("Item not found")
}
func main() {
	items := []int{1, 2, 56, 21, 43, 67, 89, 43, 21, 76}
	LinearSearch(items, 67)

}
