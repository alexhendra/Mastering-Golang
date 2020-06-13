package main

import (
	// import dengan prefix tanda titik
	// menjadikan komponen library selevel dengan file peng-import
	. "LevelAkses/library"
	"fmt"
)

func main() {
	s1 := Mahasiswa{"Alex", 29}
	fmt.Println("Nama", s1.Nama)
}
