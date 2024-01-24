package main

/*
When you write a program in Go, you will have a main package
defined with a main func inside it.
Packages are ways of grouping up related Go code together.
*/

// Importing a package which contains the Println function
import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
