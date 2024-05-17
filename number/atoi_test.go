package number

import "testing"

func TestAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				s: "-+1",
			},
			0,
		},
		{
			"",
			args{
				s: "42",
			},
			42,
		},

		{
			"",
			args{
				s: "  +42",
			},
			42,
		},

		{
			"",
			args{
				s: "-042",
			},
			-42,
		},

		{
			"",
			args{
				s: "1337c0d3",
			},
			1337,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Atoi(tt.args.s); got != tt.want {
				t.Errorf("Atoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
