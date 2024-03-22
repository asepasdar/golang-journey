package main

import "fmt"

func main() {
	const pi float32 = 3.14
	var r float32 = 7
	var luas float32 = pi * r * r

	var pangkat = 8
	pangkat *= pangkat

	var user1 string = "Asep"
	var user2 string = "Asep"
	var isSame bool = user1 == user2

	fmt.Println("Luas lingkaran adalah", luas)
	fmt.Println("Nilai dari variable pangkat", pangkat)
	fmt.Println(isSame)
}