package main

import (
	"fmt"
)

type PhysicalConstant struct {
	Value float64
	Unit  string
}

func (p PhysicalConstant) Display() {
	fmt.Println(p.Value, p.Unit)
}
