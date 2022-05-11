package helpers

import "testing"

func TestMinus(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "when happy",
			args: args{
				num1: 10,
				num2: 1,
			},
			want: 9,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Minus(test.args.num1, test.args.num2); got != test.want {
				t.Errorf("calculate.Minus() = %v, want %v", got, test.want)
			}
		})
	}
}
