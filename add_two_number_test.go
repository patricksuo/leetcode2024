package leetcode2024

import "testing"

func TestAddTowNumBitOp(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				a: 1,
				b: 1,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				a: 1,
				b: 9,
			},
			want: 10,
		},
		{
			name: "",
			args: args{
				a: 1,
				b: 3,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				a: -1,
				b: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTowNumBitOp(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AddTowNumBitOp() = %v, want %v", got, tt.want)
			}
		})
	}
}
