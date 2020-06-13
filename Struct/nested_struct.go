package main

func main() {
	// anonymous struct yang diembed ke sebuah struct
	type student struct {
		person struct {
			name string
			age  int
		}
		grade   int
		hobbies []string
	}
}
