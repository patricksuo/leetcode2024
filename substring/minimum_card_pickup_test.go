package substring

import "testing"

func TestMinimumCardPickup(t *testing.T) {
	type args struct {
		cards []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				cards: []int{1, 2, 1},
			},
			want: 3,
		},

		{
			name: "",
			args: args{
				cards: []int{1, 1, 1},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				cards: []int{1, 2},
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				cards: []int{3, 4, 2, 3, 4, 7},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				cards: []int{3, 4, 2, 3, 4, 7},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumCardPickup(tt.args.cards); got != tt.want {
				t.Errorf("MinimumCardPickup() = %v, want %v", got, tt.want)
			}
		})
	}
}
