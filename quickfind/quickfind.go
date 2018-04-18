//quickfind finds whether there is a connection between to nodes
package main

import "fmt"

type connection interface {
	connected(p, q int) bool
	union(p, q int)
}

type quickfind struct {
	data []int
}

func main() {
	fmt.Println("**** Quick find *****")
	//intialization of the array
	var qf quickfind
	qf.data = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	qf.union(0, 9)
	qf.connected(0, 9)
	qf.connected(0, 7)
	qf.union(7, 9)
	qf.connected(0, 7)

}

func (i *quickfind) connected(p, q int) bool {
	if i.data[p] == i.data[q] {
		fmt.Printf("hurrey! %d and %d are connected\n", p, q)
		return true
	}
	fmt.Printf("sorry! %d and %d are not connected\n", p, q)
	return false
}

func (i *quickfind) union(p, q int) {
	q1 := i.data[p]
	q2 := i.data[q]
	for z, k := range i.data {
		if k == q1 {
			i.data[z] = q2
		}
	}
	return
}
