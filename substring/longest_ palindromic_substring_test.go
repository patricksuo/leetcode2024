package substring

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "",
			args: args{
				s: "aba",
			},
			want: "aba",
		},

		{
			name: "",
			args: args{
				s: "aaa",
			},
			want: "aaa",
		},

		{
			name: "",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},

		{
			name: "",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestPalindromeFast(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "",
			args: args{
				s: "aba",
			},
			want: "aba",
		},

		{
			name: "",
			args: args{
				s: "aaa",
			},
			want: "aaa",
		},

		{
			name: "",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},

		{
			name: "",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindromeFast(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindromeFast() = %v, want %v", got, tt.want)
			}
		})
	}
}
