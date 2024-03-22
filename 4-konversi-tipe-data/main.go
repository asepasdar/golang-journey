package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)
	var nama string = "Asep Darmawan"

	fmt.Println(nilai32, nilai64, nilai16)
	fmt.Println("Huruf pertama adalah", string(nama[0]))
}