// When you write a program in Go you will have a main package
// defined with a main func insite it.

package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "



func Hello(name string, language string) string { 
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) { // Here we named a returned value prefix, which is a string.
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return // Returns the prefix string value. 
}




func main() {
	fmt.Println(Hello("", "English"))
	fmt.Println(Hello("Jose", "Spanish"))
}
