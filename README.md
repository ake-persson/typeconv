[![GoDoc](https://godoc.org/github.com/mickep76/cnv?status.svg)](https://godoc.org/github.com/mickep76/cnv)
[![Go Report Card](https://goreportcard.com/badge/github.com/mickep76/cnv)](https://goreportcard.com/report/github.com/mickep76/cnv)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/mickep76/cnv/blob/master/LICENSE)

# cnv

Provides conversion from a string to specific Go type

## Example

```go
package main

import (
        "fmt"
        "log"
        "strings"

        "github.com/mickep76/cnv"
)

type Person struct {
        FirstName  string
        MiddleName string
        LastName   string
}

func (p *Person) Parse(s string) error {
        a := strings.Split(s, " ")
        if len(a) > 2 {
                p.FirstName = a[0]
                p.MiddleName = a[1]
                p.LastName = a[2]
        } else if len(a) > 1 {
                p.FirstName = a[0]
                p.LastName = a[2]

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
```
