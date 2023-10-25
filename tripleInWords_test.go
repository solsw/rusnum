package rusnum

import (
	"testing"
)

func TestOnesInWords(t *testing.T) {
	type args struct {
		n      int64
		gender GrammaticalGender
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0",
			args: args{n: 0},
			want: "ноль",
		},
		{name: "1",
			args: args{n: 1},
			want: "одно",
		},
		{name: "101",
			args: args{n: 101, gender: Feminine},
			want: "сто одна",
		},
		{name: "1001",
			args: args{n: 1001, gender: Masculine},
			want: "один",
		},
		{name: "10000",
			args: args{n: 10000},
			want: "ноль",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OnesInWords(tt.args.n, tt.args.gender); got != tt.want {
				t.Errorf("OnesInWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThousandsInWords(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{n: 1},
			want: "",
		},
		{name: "101",
			args: args{n: 101},
			want: "",
		},
		{name: "1001",
			args: args{n: 1001},
			want: "одна тысяча",
		},
		{name: "12345",
			args: args{n: 12345},
			want: "двенадцать тысяч",
		},
		{name: "123456789",
			args: args{n: 123456789},
			want: "четыреста пятьдесят шесть тысяч",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThousandsInWords(tt.args.n); got != tt.want {
				t.Errorf("ThousandsInWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
