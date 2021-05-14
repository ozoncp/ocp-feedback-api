package main

import "fmt"

// Hello sends greetings to the user with a given name
func Hello(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

func main() {
	fmt.Println(Hello("Sergey Pislegin"))
}
