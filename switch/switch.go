package main

import "fmt"

func print(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Printf("The type of %v is %T\n", v, v)
	case string:
		fmt.Printf("The type of %v is %T\n", v, v)
	case bool:
		fmt.Printf("The type of %v is %T\n", v, v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	print(2)
	print("hello")
	print(2.20)
	print(true)
}
