package substring

import (
	"testing"
)

func TestMaxErasureValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 1, 1},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 1, 1},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				nums: []int{4, 2, 4, 5, 6},
			},
			want: 17,
		},
		{
			name: "",
			args: args{
				nums: []int{5, 2, 1, 2, 5, 2, 1, 2, 5},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxErasureValue(tt.args.nums); got != tt.want {
				t.Errorf("MaxErasureValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
