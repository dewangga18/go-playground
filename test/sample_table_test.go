package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test ./test -run TestTableFunction -v
func TestTableFunction(t *testing.T) {
	squareTests := []struct {
		name    string
		expect  int
		request int
	}{
		{
			name:    "Test1",
			expect:  1,
			request: 1,
		},
		{
			name:    "Test2",
			expect:  4,
			request: 2,
		},
		{
			name:    "Test3",
			expect:  9,
			request: 3,
		},
		{
			name:    "Test4",
			expect:  16,
			request: 4,
		},
		{
			name:    "Test5",
			expect:  25,
			request: 5,
		},
	}

	for _, test := range squareTests {
		t.Run(test.name, func(t *testing.T) {
			got := Square(test.request)
			assert.Equal(t, test.expect, got)
		})
	}
}
