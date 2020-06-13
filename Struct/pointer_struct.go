package main

import "fmt"

type murid struct {
	nama     string
	rangking int
}

func main() {
	var m1 = murid{nama: "Alex", rangking: 1}
	var m2 *murid = &m1

	fmt.Println("Murid 1, Nama: ", m1.nama)
	fmt.Println("Murid 2, Nama: ", m2.nama)

	m2.nama = "Joko"
	fmt.Println("Murid 1, Nama :", m1.nama)
	fmt.Println("Murid 2, Nama :", m2.nama)
}
