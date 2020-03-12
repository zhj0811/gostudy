package main

import (
	"encoding/hex"
	"fmt"
)

func singleName(a []int) int {
	var res = 0
	for i := 0; i < len(a); i++ {
		res ^= a[i]
	}
	// range i : a
	return res
}

func main() {
	// a := [...]int{2, 3, 3, 2, 9}	//不可作为slice实参
	a := []int{2, 3, 2, 3, 4}
	n := singleName(a)
	fmt.Println("single number:", n)

	var c byte = 'o'
	fmt.Printf("%s\n", string(c))
	var d byte
	const hextable = "0123456789abcdef"
	fmt.Printf("c: %b\n", c)
	fmt.Printf("c>>4: %b\n", c>>4)
	fmt.Printf("c>>4.T: %T\n", c>>4)
	d = hextable[c>>4] //hextable[6]
	fmt.Printf("%d\n", d)
	fmt.Printf("%x\n", d)
	fmt.Printf("%c\n", d)
	// fmt.Printf("%+v\n", d)
	// fmt.Printf("%s\n", string(d))
	e := hex.EncodeToString([]byte{c})
	fmt.Printf("e: %s\n", e)
}
