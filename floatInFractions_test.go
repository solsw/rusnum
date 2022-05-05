package rusnum

import (
	"fmt"
	"math"
	"testing"
)

func TestFloatInFractions(t *testing.T) {
	type args struct {
		f        float64
		fraction Fraction
		showZero bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0.0_1",
			args: args{
				f:        0.0,
				fraction: Tenth,
				showZero: false,
			},
			want: "",
		},
		{name: "0.0_2",
			args: args{
				f:        0.0,
				fraction: Tenth,
				showZero: true,
			},
			want: "ноль десятых",
		},
		{name: "0.0_3",
			args: args{
				f:        0.0,
				fraction: Milliardth,
				showZero: true,
			},
			want: "ноль миллиардных",
		},
		{name: "0.007",
			args: args{
				f:        0.007,
				fraction: Hundredth,
			},
			want: "",
		},
		{name: "1.0",
			args: args{
				f:        1.0,
				fraction: Tenth,
			},
			want: "десять десятых",
		},
		{name: "1.1",
			args: args{
				f:        1.1,
				fraction: Tenth,
			},
			want: "одиннадцать десятых",
		},
		{name: "1.29_10",
			args: args{
				f:        1.29,
				fraction: Tenth,
			},
			want: "двенадцать десятых",
		},
		{name: "1.234_10",
			args: args{
				f:        1.234,
				fraction: Tenth,
			},
			want: "двенадцать десятых",
		},
		{name: "1.234_100",
			args: args{
				f:        1.234,
				fraction: Hundredth,
			},
			want: "сто двадцать три сотых",
		},
		{name: "1.234_1000",
			args: args{
				f:        1.234,
				fraction: Thousandth,
			},
			want: "одна тысяча двести тридцать четыре тысячных",
		},
		{name: "0.8",
			args: args{
				f:        0.8,
				fraction: Tenth,
			},
			want: "восемь десятых",
		},
		{name: "-0.8_10",
			args: args{
				f:        -0.8,
				fraction: Tenth,
			},
			want: "минус восемь десятых",
		},
		{name: "-0.8_1000",
			args: args{
				f:        -0.8,
				fraction: Thousandth,
			},
			want: "минус восемьсот тысячных",
		},
		{name: "-0.8_1000000",
			args: args{
				f:        -0.8,
				fraction: Millionth,
			},
			want: "минус восемьсот тысяч миллионных",
		},
		{name: "0.1000000001_1000000000",
			args: args{
				f:        0.1000000001,
				fraction: Tenmilliardth,
			},
			want: "один миллиард одна десятимиллиардная",
		},
		{name: "0.12345678910_10000000000",
			args: args{
				f:        0.1234567891,
				fraction: Hundredmilliardth,
			},
			want: "двенадцать миллиардов триста сорок пять миллионов шестьсот семьдесят восемь тысяч девятьсот десять стомиллиардных",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FloatInFractions(tt.args.f, tt.args.fraction, tt.args.showZero)
			if got != tt.want {
				t.Errorf("FloatInFractions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleFloatInFractions() {
	fmt.Println(FloatInFractions(21, Tenth, false))
	fmt.Println(FloatInFractions(math.Ln2, Tenmilliardth, false))
	// Output:
	// двести десять десятых
	// шесть миллиардов девятьсот тридцать один миллион четыреста семьдесят одна тысяча восемьсот пять десятимиллиардных
}

func ExampleFloatInFractionsAuto() {
	fmt.Println(FloatInFractionsAuto(21, false))
	fmt.Println(FloatInFractionsAuto(12.34, false))
	fmt.Println(FloatInFractionsAuto(0.123456789, false))
	fmt.Println(FloatInFractionsAuto(math.Log10E, false))
	// Output:
	// двадцать один
	// одна тысяча двести тридцать четыре сотых
	// сто двадцать три миллиона четыреста пятьдесят шесть тысяч семьсот восемьдесят девять миллиардных
	// сорок три миллиарда четыреста двадцать девять миллионов четыреста сорок восемь тысяч сто девянoсто стомиллиардных
}