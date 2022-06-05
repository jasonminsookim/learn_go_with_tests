// When you write a program in Go you will have a main package
// defined with a main func insite it. 

package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name 
}
func main() {
	fmt.Println(Hello("Jason"))
}
