package main

import "fmt"

func main() {
	fmt.Println("\n To find if there are pairs in an array which sum to a value x ")

	input := [6]int{2, 3, 6, 9, 8, 7}
	val := 10
	m := make(map[int]int)
	for _, k := range input {
		if k < val {
			m[val-k] = k
		}
		if val, ok := m[k]; ok {
			fmt.Println("The pair are ", val, " and ", k)
		}
	}
}
