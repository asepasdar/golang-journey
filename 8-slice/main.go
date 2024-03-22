package main

import "fmt"

func main() {
	var names = [...]string{
		"Asep",
		"Riki",
		"Marco",
		"Michael",
		"David",
	}

	slice := names[1:3]
	fmt.Println(slice)

	slice = names[:3]
	fmt.Println(slice)

	slice = names[3:]
	fmt.Println(slice)

	slice = names[:]
	fmt.Println(slice)
}