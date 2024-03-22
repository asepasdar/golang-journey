package main

import "fmt"

func main() {
	var float float32 = 3.14
	var number uint8 = 9
	var number2 int16 = -9
	var condition bool = true
	var (
		name string = "Asep Darmawan"
		jenisKelamin string = "Laki - laki"
	)

	fmt.Println("Nilai dari float adalah ", float)
	fmt.Println("Nilai dari number adalah ", number)
	fmt.Println("Nilai dari number2 adalah ", number2)
	fmt.Println("Nilai dari condition adalah ", condition)
	fmt.Println("Nilai dari name adalah ", name)
	fmt.Println("Jenis Kelamin adalah ", jenisKelamin)
	fmt.Println("Panjang dari name adalah ", len(name))
	fmt.Println("Karakter Awal dari name adalah ", name[0])
	fmt.Println("Karakter Terakhir dari name adalah ", name[len(name)-1])
}