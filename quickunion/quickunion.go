//quick union is an improvemnet for quick find for union operation
package main

import "fmt"

type connection interface {
	connected(p, q int) bool
	union(p, q int)
	root(p int) int
}
type quickunion struct {
	data []int
}

func main() {
	fmt.Println("**** Quick Union ****")

	//intialization of the array
	var qu quickunion
	qu.data = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	qu.union(0, 9)
	qu.union(1, 2)
	qu.union(5, 6)
	qu.union(6, 9)
	fmt.Println(qu.data)
}

func (u *quickunion) connected(p, q int) bool {
	return u.root(p) == u.root(q)
}

func (u *quickunion) root(p int) int {
	for p != u.data[p] {
		p = u.data[p]
	}
	return p
}

func (u *quickunion) union(p, q int) {
	u.data[u.root(p)] = u.root(q)
	return
}
