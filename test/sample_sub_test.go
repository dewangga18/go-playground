package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubFunction(t *testing.T) {
	t.Run("sub1", func(t *testing.T) {
		res := Square(2)

		assert.Equal(t, 4, res)
	})
	t.Run("sub2", func(t *testing.T) {
		res := Square(3)

		assert.Equal(t, 9, res)
	})
}