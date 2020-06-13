package main

import "fmt"

func main() {
	// slice yang tipe datanya anonymous struct
	// dan langsung diberikan nilai inisialisasi
	var allStudents = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}

	for _, student := range allStudents {
		fmt.Println(student)
	}
}
