package fib_test

import (
	"../fib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFib(t *testing.T) {
	testCases := []struct{
		name string
		in int
		out int
	}{
		{
			name: "zero",
			in:  0,
			out: 0,
		},
		{
			name: "one",
			in: 1,
			out: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.out, fib.Fib(tc.in))
		})
	}
}