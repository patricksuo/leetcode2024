package leetcode2024

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				[]int{1, 2, 3},
				3,
			},
			want: []int{0, 1},
		},

		{
			name: "",
			args: args{
				[]int{1, 1, 3},
				2,
			},
			want: []int{0, 1},
		},
		{
			name: "",
			args: args{
				[]int{2, 7, 11, 15},
				9,
			},
			want: []int{0, 1},
		},
		{
			name: "",
			args: args{
				[]int{3, 2, 4},
				6,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.array, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSumSortedArray(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				[]int{1, 2, 3},
				3,
			},
			want: []int{0, 1},
		},

		{
			name: "",
			args: args{
				[]int{1, 1, 3},
				2,
			},
			want: []int{0, 1},
		},
		{
			name: "",
			args: args{
				[]int{2, 7, 11, 15},
				9,
			},
			want: []int{0, 1},
		},
		{
			name: "",
			args: args{
				[]int{3, 2, 4},
				6,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSumSortedArray(tt.args.array, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
