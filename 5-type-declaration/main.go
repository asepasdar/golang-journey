package main

import "fmt"

func main() {
	type NoKtp string

	var myKtp NoKtp = "107200027938123"
	var otherKtp string = "10029351238850123"

	fmt.Println(myKtp)
	fmt.Println(NoKtp(otherKtp))
}