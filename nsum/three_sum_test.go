package nsum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{
			name: "",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			wantResult: [][]int{
				{-1, 0, 1},
				{-1, -1, 2},
			},
		},

		{
			name: "",
			args: args{
				nums: []int{0, 1, 1},
			},
			wantResult: nil,
		},

		{
			name: "",
			args: args{
				nums: []int{0, 0, 0},
			},
			wantResult: [][]int{
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ThreeSum(tt.args.nums); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ThreeSum() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
