package bitoperation

import (
	"testing"
)

func TestReverseIntegerx(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "",
			args: args{
				x: -123,
			},
			want: -321,
		},
		{
			name: "",
			args: args{
				x: 120,
			},
			want: 21,
		},
		{
			name: "",
			args: args{
				x: 1534236469,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseIntegerx(tt.args.x); got != tt.want {
				t.Errorf("ReverseIntegerx() = %v, want %v", got, tt.want)
			}
		})
	}
}
