package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/mickep76/cnv"
)

func main() {
	i, err := cnv.Conv("123", reflect.Int8)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("int8: %v\n", i)

	b, err := cnv.Conv("true", reflect.Bool)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("bool: %v\n", b)
}
