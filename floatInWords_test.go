package rusnum

import (
	"fmt"
	"math"
	"testing"
)

func TestFloatInWords(t *testing.T) {
	type args struct {
		f         float64
		fraction  Fraction
		binder    Binder
		zeroInt   bool
		zeroFrac  bool
		withZeros bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "01",
			args: args{
				f: 0.0,
			},
			want: "ноль",
		},
		{name: "02",
			args: args{f: 0.0,
				fraction: Tenth,
				zeroFrac: true,
			},
			want: "ноль десятых",
		},
		{name: "03",
			args: args{
				f:        0.0,
				fraction: Tenth,
				zeroInt:  true,
				zeroFrac: true,
			},
			want: "ноль ноль десятых",
		},
		{name: "04",
			args: args{
				f:        -0.0,
				fraction: Tenth,
				binder:   And,
				zeroInt:  true,
				zeroFrac: true,
			},
			want: "ноль и ноль десятых",
		},
		{name: "05",
			args: args{
				f:        0.0,
				fraction: Tenth,
				binder:   Whole,
				zeroInt:  true,
				zeroFrac: true,
			},
			want: "ноль целых ноль десятых",
		},
		{name: "06",
			args: args{
				f:        0.0,
				fraction: Thousandth,
				zeroFrac: true,
			},
			want: "ноль тысячных",
		},
		{name: "0.1_10",
			args: args{
				f:        0.1,
				fraction: Tenth,
				zeroFrac: true,
			},
			want: "одна десятая",
		},
		{name: "-0.1_100",
			args: args{
				f:        -0.1,
				fraction: Hundredth,
				zeroFrac: true,
			},
			want: "минус десять сотых",
		},
		{name: "0.1_1000",
			args: args{
				f:        0.1,
				fraction: Thousandth,
				zeroFrac: true,
			},
			want: "сто тысячных",
		},
		{name: "21.123_100_Whole",
			args: args{
				f:        21.123,
				fraction: Hundredth,
				binder:   Whole,
				zeroFrac: true,
			},
			want: "двадцать одна целая двенадцать сотых",
		},
		{name: "-21.123_1000000000_And",
			args: args{
				f:        -21.123,
				fraction: Milliardth,
				binder:   And,
				zeroFrac: true,
			},
			want: "минус двадцать один и сто двадцать три миллиона миллиардных",
		},
		{name: "2.3",
			args: args{
				f:        2.3,
				fraction: Tenth,
				binder:   And,
			},
			want: "два и три десятых",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FloatInWords(tt.args.f, tt.args.fraction, tt.args.binder, tt.args.zeroInt, tt.args.zeroFrac, tt.args.withZeros)
			if got != tt.want {
				t.Errorf("FloatInWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleFloatInWords() {
	fmt.Println(FloatInWords(21, Tenth, Whole, false, true, false))
	fmt.Println(FloatInWords(math.E, Hundredth, NoBinder, false, true, false))
	fmt.Println(FloatInWords(math.E, Thousandth, And, false, true, false))
	fmt.Println(FloatInWords(math.E, Millionth, Whole, false, true, false))
	// Output:
	// двадцать одна целая ноль десятых
	// два семьдесят одна сотая
	// два и семьсот восемнадцать тысячных
	// две целых семьсот восемнадцать тысяч двести восемьдесят одна миллионная
}

func ExampleFloatInWordsAuto() {
	fmt.Println(FloatInWordsAuto(21, Whole, false, false, false))
	fmt.Println(FloatInWordsAuto(math.Pi, Whole, false, false, false))
	// Output:
	// двадцать одна целая
	// три целых четырнадцать миллиардов сто пятьдесят девять миллионов двести шестьдесят пять тысяч триста пятьдесят восемь стомиллиардных
}

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
