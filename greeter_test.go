package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGreeter_Greet(t *testing.T) {
	g := Greeter{Name: "World"}
	result := g.Greet()
	require.NotEmpty(t, result, "Greet() should not return an empty string")
	assert.Equal(t, "Hello, World!", result)
}

func TestGreeter_Greet_EmptyName(t *testing.T) {
	g := Greeter{Name: ""}
	result := g.Greet()
	assert.Equal(t, "Hello, !", result)
}

func TestGreeter_Greet_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple name", "Alice", "Hello, Alice!"},
		{"another name", "Bob", "Hello, Bob!"},
		{"empty name", "", "Hello, !"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := Greeter{Name: tc.input}
			assert.Equal(t, tc.expected, g.Greet())
		})
	}
}