package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var firstname, lastname string
	fmt.Scan(&firstname, &lastname)

	fmt.Printf("Hello %s %s, welcome to the first GO program", firstname, lastname)
}
