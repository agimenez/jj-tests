// Hello world program
package main

import "fmt"

// add documentation for main
func main() {
	p("Hello World!")
	p("Goodbye World!")
}

// a function that prints a message
func p(s string) {
	fmt.Println(s)
}
