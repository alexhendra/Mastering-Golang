package main

import "LevelAkses/library"

func main() {
	library.SayHello()
	library.introduce("Alex") // error karena mengakses private function (Unexported)
}