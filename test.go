package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	var testing = [2]string{"Hello", "world"}
	fmt.Println(testing[0] + " " + testing[1])

	if 1 > 0 {
		fmt.Println("Hello")
	} else {
		fmt.Println("World")
	}

	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	}
}
