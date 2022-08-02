package main

import "fmt"

func main() {

	s := make([]string, 2)
	s[0] = "dog"
	s[1] = "cat"
	fmt.Println("set:", s)
	fmt.Println("len:", len(s))

}
