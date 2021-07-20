package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ake-persson/cnv"
)

// Person struct.
type Person struct {
	FirstName  string
	MiddleName string
	LastName   string
}

// Parse person string to struct.
func (p *Person) Parse(s string) error {
	a := strings.Split(s, " ")
	if len(a) > 2 {
		p.FirstName = a[0]
		p.MiddleName = a[1]
		p.LastName = a[2]
	} else if len(a) > 1 {
		p.FirstName = a[0]
		p.LastName = a[1]
	} else {
		p.FirstName = a[0]
	}
	return nil
}

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

	var p Person
	if err := cnv.Parse("John Sam Doe", &p); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("First: %s Middle: %s Last: %s\n", p.FirstName, p.MiddleName, p.LastName)
}
