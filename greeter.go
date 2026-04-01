package main

import "fmt"

// Greeter holds greeting configuration.
type Greeter struct {
	Name string
}

// Greet returns a greeting string for the configured name.
func (g Greeter) Greet() string {
	return fmt.Sprintf("Hello, %s!", g.Name)
}