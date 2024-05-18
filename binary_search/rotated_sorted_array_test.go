package binarysearch

import "testing"

func TestFindInRotatedSortedArray(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				[]int{4, 5, 6, 7, 0, 1, 2},
				4,
			},
			0,
		},

		{
			"",
			args{
				[]int{4, 5, 6, 7, 0, 1, 2},
				3,
			},
			-1,
		},

		{
			"",
			args{
				[]int{4},
				0,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindInRotatedSortedArray(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("FindInRotatedSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
