package example

// Package example is a package for showing how examples can be tests and documentation
import (
	"fmt"
)

type Language interface {
	Greet(string) string
}

type English struct{}

// Greet prints hello to the name provided
func (e English) Greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

type Spanish struct{}

func (s Spanish) Greet(name string) string {
	return fmt.Sprintf("Hola, %s", name)
}

// Page will print out a message asking each person who hasn't checked in to do so.
func Page(checkIns map[string]bool) {
	for name, checkIn := range checkIns {
		if !checkIn {
			fmt.Printf("Paging %s; please see the front desk to check in.", name)
		}
	}
}
