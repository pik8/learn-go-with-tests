package main

/*
When you write a program in Go, you will have a main package
defined with a main func inside it.
Packages are ways of grouping up related Go code together.
*/

// Importing a package which contains the Println function
import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
