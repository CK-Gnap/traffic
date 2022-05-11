package helpers

import "testing"

func TestRandFloat(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "when happy get between 1 and 3",
			args: args{
				min: 1,
				max: 3,
			},
			want: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := RandInt(test.args.min, test.args.max)
			if got < 1 && got > 3 {
				t.Errorf("randInt() = %v, want %v", got, test.want)
			}
		})
	}
}
