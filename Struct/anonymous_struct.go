package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// deklarasi objek dari struct anonymous
	var s1 = struct {
		person
		grade int
	}{}
	s1.person = person{"wick", 21}
	s1.grade = 2

	fmt.Println("name  :", s1.person.name)
	fmt.Println("age   :", s1.person.age)
	fmt.Println("grade :", s1.grade)

	// deklarasi object dengan var
	var student struct {
		person
		grade int
	}

	student.person = person{"wick", 21}
	student.grade = 2
}
