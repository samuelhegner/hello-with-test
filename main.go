package main

import "fmt"

func main() {
	fmt.Println(Hello("Samuel"))
}

func Hello(name string) string {
	return "Hello " + name + "!"
}
