package rusnum

import (
	"math"
	"reflect"
	"testing"
)

func Test_fractionFromFloat(t *testing.T) {
	type args struct {
		fl float64
	}
	tests := []struct {
		name string
		args args
		want Fraction
	}{
		{name: "0", args: args{fl: 123456789}, want: NoFraction},
		{name: "1", args: args{fl: 12345678.9}, want: Tenth},
		{name: "2", args: args{fl: 1234567.89}, want: Hundredth},
		{name: "3", args: args{fl: 123456.789}, want: Thousandth},
		{name: "4", args: args{fl: 12345.6789}, want: Tenthousandth},
		{name: "5", args: args{fl: 1234.56789}, want: Hundredthousandth},
		{name: "6", args: args{fl: 123.456789}, want: Millionth},
		{name: "7", args: args{fl: 12.3456789}, want: Tenmillionth},
		{name: "8", args: args{fl: 1.23456789}, want: Hundredmillionth},
		{name: "9", args: args{fl: .123456789}, want: Milliardth},
		{name: "E", args: args{fl: math.E}, want: Hundredmilliardth},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionFromFloat(tt.args.fl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fractionFromFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
