package rusnum

import (
	"fmt"
	"math"
	"testing"
)

func TestIntInWords(t *testing.T) {
	type args struct {
		n         int64
		withZeros bool
		gender    GrammaticalGender
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0",
			args: args{
				n: 0,
			},
			want: "ноль",
		},
		{name: "0tm",
			args: args{
				n:         0,
				withZeros: true,
				gender:    Masculine,
			},
			want: "ноль",
		},
		{name: "-1tm",
			args: args{
				n:         -1,
				withZeros: true,
				gender:    Masculine,
			},
			want: "минус один",
		},
		{name: "1tf",
			args: args{
				n:         1,
				withZeros: true,
				gender:    Feminine,
			},
			want: "одна",
		},
		{name: "100ff",
			args: args{
				n:         100,
				withZeros: false,
				gender:    Feminine,
			},
			want: "сто",
		},
		{name: "1000tm",
			args: args{
				n:         1000,
				withZeros: true,
				gender:    Masculine,
			},
			want: "одна тысяча ноль",
		},
		{name: "1000fm",
			args: args{
				n:         1000,
				withZeros: false,
				gender:    Masculine,
			},
			want: "одна тысяча",
		},
		{name: "1001fm",
			args: args{
				n:         1001,
				withZeros: false,
				gender:    Masculine,
			},
			want: "одна тысяча один",
		},
		{name: "2000001",
			args: args{
				n: 2000001,
			},
			want: "два миллиона одно",
		},
		{name: "2000002ft",
			args: args{
				n:         2000002,
				withZeros: true,
				gender:    Feminine,
			},
			want: "два миллиона ноль тысяч две",
		},
		{name: "2000000000000002ff",
			args: args{
				n:         2000000000000002,
				withZeros: false,
				gender:    Feminine,
			},
			want: "два квадриллиона две",
		},
		{name: "2000000000000002nt",
			args: args{
				n:         2000000000000002,
				withZeros: true,
				gender:    Neuter,
			},
			want: "два квадриллиона ноль триллионов ноль миллиардов ноль миллионов ноль тысяч два",
		},
		{name: "math.MinInt64",
			args: args{
				n: math.MinInt64,
			},
			want: minInt64InWords,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntInWords(tt.args.n, tt.args.withZeros, tt.args.gender); got != tt.want {
				t.Errorf("IntInWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIntInWords() {
	fmt.Println(IntInWords(21, true, Neuter))
	fmt.Println(IntInWords(-1032, true, Feminine))
	fmt.Println(IntInWords(2000001, true, Feminine))
	fmt.Println(IntInWords(2000000001, false, Masculine))
	// Output:
	// двадцать одно
	// минус одна тысяча тридцать две
	// два миллиона ноль тысяч одна
	// два миллиарда один
}
