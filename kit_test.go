package gokit

import "testing"

func TestPrintJSON(t *testing.T) {
	c := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  30,
	}

	PrintJSON(c)
}
