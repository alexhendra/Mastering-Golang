package main

import "fmt"

type manusia struct {
	nama string
	umur int
}

// Murid memiliki struct di dalam struct (embedded struct)
type murid struct {
	rangking int
	manusia
}

func main() {
	var m1 = murid{}
	m1.nama = "Alex"
	m1.umur = 21
	m1.rangking = 2

	fmt.Println("Name: ", m1.nama)
	fmt.Println("Umur: ", m1.umur) // bisa langsung memanggil property dari struct parent tanpa harus menulis nama parentnya
	fmt.Println("Umur: ", m1.manusia.umur)

	var mn = manusia{nama: "Joko", umur: 30}
	var m2 = murid{rangking: 1, manusia: mn}

	fmt.Println("m2 Name: ", m2.nama)
	fmt.Println("m2 Umur: ", m2.umur) // bisa langsung memanggil property dari struct parent tanpa harus menulis nama parentnya
	fmt.Println("m2 Rangking: ", m2.rangking)
}
