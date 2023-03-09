package rusnum

import (
	"math"
	"reflect"
	"testing"
)

func Test_triplesInWords(t *testing.T) {
	type args struct {
		n         int64
		withZeros bool
		gender    GrammaticalGender
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "01",
			args: args{n: 0, withZeros: false},
			want: []string{"ноль"},
		},
		{name: "02",
			args: args{n: 0, withZeros: true},
			want: []string{"ноль"},
		},
		{name: "1n",
			args: args{n: 1, gender: Neuter},
			want: []string{"одно"},
		},
		{name: "2f",
			args: args{n: 2, gender: Feminine},
			want: []string{"две"},
		},
		{name: "21m",
			args: args{n: 21, gender: Masculine},
			want: []string{"двадцать один"},
		},
		{name: "1000001",
			args: args{n: 1000001, withZeros: false, gender: Masculine},
			want: []string{"один", "", "один миллион"},
		},
		{name: "1000001z",
			args: args{n: 1000001, withZeros: true, gender: Masculine},
			want: []string{"один", "ноль тысяч", "один миллион"},
		},
		{name: "1001001",
			args: args{n: 1001001, gender: Neuter},
			want: []string{"одно", "одна тысяча", "один миллион"},
		},
		{name: "MaxInt64",
			args: args{n: math.MaxInt64, gender: Feminine},
			want: []string{"восемьсот семь", "семьсот семьдесят пять тысяч", "восемьсот пятьдесят четыре миллиона",
				"тридцать шесть миллиардов", "триста семьдесят два триллиона", "двести двадцать три квадриллиона", "девять квинтиллионов"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triplesInWords(tt.args.n, tt.args.withZeros, tt.args.gender); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("triplesInWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
