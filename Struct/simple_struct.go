package main

import "fmt"

type mahasiswa struct {
	nama     string
	rangking int
}

func main() {
	var s1 mahasiswa
	s1.nama = "Alex Hendra"
	s1.rangking = 5

	s2 := mahasiswa{"Budi", 1}

	var s3 = mahasiswa{nama: "jason"}

	fmt.Println("Nama: ", s1.nama)
	fmt.Println("Rangking: ", s1.rangking)

	fmt.Println("Nama: ", s2.nama)
	fmt.Println("Rangking: ", s2.rangking)

	fmt.Println("Nama: ", s3.nama)
	fmt.Println("Rangking: ", s3.rangking)
}
