package substring

import "testing"

func TestCountCompleteSubarrays(t *testing.T) {
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
				nums: []int{3, 1, 5, 11, 13},
			},
			want: 1,
		},

		{
			name: "",
			args: args{
				nums: []int{1, 1},
			},
			want: 3,
		},

		{
			name: "",
			args: args{
				nums: []int{1, 3, 1, 2, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountCompleteSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("CountCompleteSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
