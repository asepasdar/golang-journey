package main

import "fmt"

func main() {
	const pi float32 = 3.14
	const (
		firstName string = "Asep"
		lastName string = "Darmawan"
		umur uint8 = 27
	)


	fmt.Println("Nilai pi adalah", pi)
	fmt.Println(firstName, lastName, umur)
}