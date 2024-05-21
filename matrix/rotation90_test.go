package matrix

import "testing"

func TestRotateNightDegree(t *testing.T) {
	tests := []struct {
		matrix [][]int
	}{

		{
			[][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Logf("before: %v", tt.matrix)
			RotateNightDegree(tt.matrix)
			t.Logf("after %v", tt.matrix)
		})
	}
}
