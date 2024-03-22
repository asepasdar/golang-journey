package main

import "fmt"

func main() {
	var names [3]string
	var scores = [3]uint8{
		80,
		90,
	}
	var status = [...]string{
		"Pass",
		"Pass",
		"Failed",
	}

	names[0] = "asep"
	names[1] = "darmawan"
	names[2] = "michael"

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i], scores[i], " : ", status[i])
	}
}