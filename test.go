package main

import (
	"fmt"
)

func main() {
	var first, second int = 12, 42
	var p *int

	p = &first
	fmt.Println("p is", *p)
	p = &second
	fmt.Println("p is now", *p)
	*p = 1024
	fmt.Println("and now p =", *p)

}
