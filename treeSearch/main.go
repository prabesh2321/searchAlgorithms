package main

import (
	"fmt"
)

func main() {
	s := [4]int{0, 0, 0, 1}
	e := [4]int{0, 2, 1, 1}

	status := nodeDistant(4, s, e, 0)
	if !status {
		fmt.Println("Sorry there are no nodes distant at ", 4, " from ", 0)
	}

}

func nodeDistant(k int, s, e [4]int, node int) bool {

	for i := 0; i < len(s); i++ {
		if s[i] == node {
			if k == e[i] {
				fmt.Println("The nodes are ", i)
				return true
			} else if k > e[i] && e[i] != 0 {
				return false || nodeDistant(k-e[i], s, e, i)
			}
		}
	}
	return false
}
