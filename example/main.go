package main

import (
	"fmt"
	"log"

	"github.com/mickep76/cnv"
)

func main() {
	var i int8
	if err := cnv.Parse("123", &i); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("int8: %v\n", i)

	var b bool
	if err := cnv.Parse("true", &b); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("bool: %v\n", b)
}
