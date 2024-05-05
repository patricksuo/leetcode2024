package substring

import "testing"

func TestOptimalPartitionString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: "aba",
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				s: "abab",
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				s: "abacaba",
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				s: "ssssss",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptimalPartitionString(tt.args.s); got != tt.want {
				t.Errorf("OptimalPartitionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
