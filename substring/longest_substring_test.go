package substring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
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
				s: "aa",
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
				s: "abbcdaa",
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				s: "abbca",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				s: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
