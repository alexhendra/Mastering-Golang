package library

import "fmt"

// SayHello Public function dari package library
func SayHello() {
	fmt.Println("hello")
}

func introduce(name string) {
	fmt.Println("nama saya", name)
}

type Mahasiswa struct {
	Nama     string
	rangking int
}

// Student adalah anonymous struct. harus ada inisialisasi walaupun cuma kurung kurawal kosong
var Student = struct {
	Name  string
	Grade int
}{}

// fungsi init akan dieksekusi pertama kali pada file peng-import
func init() {
	Student.Name = "Alex Hendra"
	Student.Grade = 5

	fmt.Println("--> library/library.go diimport")
}
