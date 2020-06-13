package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int
	count2 := new(int)
	countTmp := 5
	count3 := &countTmp

	t := &time.Timer{}
	fmt.Printf("alamat memory count1: %#v\n", count1)
	fmt.Printf("alamat memory count2: %#v\n", count2)
	fmt.Printf("alamat memory count3: %#v\n", count3)
	fmt.Printf("alamat memory time : %#v\n", t)

	fmt.Printf("value count3: %#v\n", *count3)
	fmt.Printf("value time : %#v\n", *t)
}
