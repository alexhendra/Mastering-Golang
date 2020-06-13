package main

func main() {
	// Diberikan tag, yang bsa digunakan untuk decode/encode JSON,
	// Tag ini bisa diakses juga dari reflect
	type person struct {
		name string `tag1`
		age  int    `tag2`
	}
}
