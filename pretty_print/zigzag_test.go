package prettyprint

import "testing"

func TestZigzagView(t *testing.T) {
	type args struct {
		s   string
		row int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s:   "aba",
				row: 1,
			},
			want: "aba",
		},

		{
			name: "",
			args: args{
				s:   "PAYPALISHIRING",
				row: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},

		{
			name: "",
			args: args{
				s:   "PAYPALISHIRING",
				row: 4,
			},
			want: "PINALSIGYAHRPI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZigzagView(tt.args.s, tt.args.row); got != tt.want {
				t.Errorf("ZigzagView() = %v, want %v", got, tt.want)
			}
		})
	}
}
