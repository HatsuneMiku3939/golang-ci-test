package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		A int
		B int
		S int
	}{
		{A: 3, B: 9, S: 12},
		{A: 9, B: 3, S: 12},
	}

	for _, c := range cases {
		s := Add(c.A, c.B)
		assert.Equalf(t, c.S, s, "Add(%d, %d) should be %d, but %d", c.A, c.B, c.S, s)
	}
}
