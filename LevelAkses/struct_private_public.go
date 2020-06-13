package main

import (
	"LevelAkses/library"
	"fmt"
)

func main() {
	var s1 = library.Mahasiswa{"Alex", 25}
	fmt.Println("Nama", s1.Nama)
	fmt.Println("grade", s1.rangking) // Error: karena mengakses method private (Unexported Method)
}
