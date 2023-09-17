package addupto

import "testing"

var tests = []struct {
	name string
	arg  int
	want int
}{
	{name: "zero input",
		arg:  0,
		want: 0,
	},
	{name: "negative input",
		arg:  -1,
		want: 0,
	},
	{name: "add up to 25",
		arg:  25,
		want: 325,
	},
	{name: "add up to 1000",
		arg:  1000,
		want: 500500,
	},
}

func TestAddUptoA(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := addUptoA(tc.arg); got != tc.want {
				t.Errorf("got: %d; want: %d", got, tc.want)
			}
		})
	}
}
func TestAddUptoB(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := addUptoB(tc.arg); got != tc.want {
				t.Errorf("got: %d; want: %d", got, tc.want)
			}
		})
	}
}

func BenchmarkAddUptoA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addUptoA(10000000000)
	}
}
func BenchmarkAddUptoB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addUptoB(10000000000)
	}
}
