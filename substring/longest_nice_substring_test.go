package substring

import "testing"

func TestLongestNiceSubarray(t *testing.T) {
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
				nums: []int{1},
			},
			want: 1,
		},
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
				nums: []int{1, 2, 1, 2},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 3, 8, 48, 10},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				nums: []int{3, 1, 5, 11, 13},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{178830999, 19325904, 844110858, 806734874, 280746028, 64, 256, 33554432, 882197187, 104359873, 453049214, 820924081, 624788281, 710612132, 839991691},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestNiceSubarray(tt.args.nums); got != tt.want {
				t.Errorf("LongestNiceSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
