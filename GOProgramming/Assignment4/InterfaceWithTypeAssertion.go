package main

import "fmt"

func main() {
	var i interface{}
	i = "Hello Go"

	str := i.(string) // type assertion

	fmt.Println("Value:", str)
}
