package test

import "testing"

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(10)
	}
}

func BenchmarkSquareSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Run("Square 5", func(b *testing.B) {
			Square(5)
		})
		b.Run("Square 3", func(b *testing.B) {
			Square(3)
		})
	}
}

func BenchmarkTableSquare(b *testing.B) {
	testCase := []struct {
		Name     string
		Input    int
		Expected int
	}{
		{
			Name:     "Square 5",
			Input:    5,
			Expected: 25,
		},
		{
			Name:     "Square 3",
			Input:    3,
			Expected: 9,
		},
		{
			Name:     "Square 2",
			Input:    2,
			Expected: 4,
		},
		{
			Name:     "Square 1",
			Input:    1,
			Expected: 1,
		},
	}

	for _, tc := range testCase {
		b.Run(tc.Name, func(b *testing.B) {
			Square(tc.Input)
		})
	}
}
