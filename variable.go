package main

import "fmt"

func main() {
	name := "Andika"
	age := 29
	WNI := true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(WNI)

	// membuat variabel (var) /constant (const) sekaligus banyak
	var (
		firstName = "Joko"
		lastName  = "Purwanto"
	)

	const (
		NamaAwal  = "Joko"
		NamaAkhir = "Purwanto"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
