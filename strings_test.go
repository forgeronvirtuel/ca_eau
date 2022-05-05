package ca_eau

import (
	"reflect"
	"testing"
)

func TestIsSubstringOf(t *testing.T) {
	var txt, s string
	var got, exp bool

	txt, s = "Bonjour", "jour"
	got, exp = IsSubstringOf(txt, s), true
	if got != exp {
		t.Errorf("IsSubstringOf(%s, %s) = %t, expected = %t", txt, s, got, exp)
	}

	txt, s = "Bonjour", "t"
	got, exp = IsSubstringOf(txt, s), false
	if got != exp {
		t.Errorf("IsSubstringOf(%s, %s) = %t, expected = %t", txt, s, got, exp)
	}

	txt, s = "Bonjour", "onj"
	got, exp = IsSubstringOf(txt, s), true
	if got != exp {
		t.Errorf("IsSubstringOf(%s, %s) = %t, expected = %t", txt, s, got, exp)
	}

	txt, s = "Bonjour", "Bonj"
	got, exp = IsSubstringOf(txt, s), true
	if got != exp {
		t.Errorf("IsSubstringOf(%s, %s) = %t, expected = %t", txt, s, got, exp)
	}

	txt, s = "Bonjour", ""
	got, exp = IsSubstringOf(txt, s), true
	if got != exp {
		t.Errorf("IsSubstringOf(%s, %s) = %t, expected = %t", txt, s, got, exp)
	}

	txt, s = "Bonjour", "Bon"
	got, exp = IsSubstringOf(txt, s), true
	if got != exp {
		t.Errorf("IsSubstringOf(%s, %s) = %t, expected = %t", txt, s, got, exp)
	}
}

func TestUpperOneOnTwo(t *testing.T) {
	var v, got, exp string

	v, exp = "Hello world !", "HeLlO wOrLd !"
	got = UpperOneOnTwo(v)
	if got != exp {
		t.Errorf("UpperOneOnTwo(%s) = %s, expected = %s", v, got, exp)
	}
}

func TestUpperFirstCharEachWord(t *testing.T) {
	var v, got, exp string

	v, exp = "this is a text", "This Is A Text"
	got = UpperFirstCharEachWord(v)
	if got != exp {
		t.Errorf("UpperFirstCharEachWord(%s) = %s, exp = %s", v, got, exp)
	}
}

func TestIsOnlyDigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal case",
			args: args{s: "1726371361"},
			want: true,
		},
		{
			name: "error case",
			args: args{s: "82731 Bonjour"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOnlyDigit(tt.args.s); got != tt.want {
				t.Errorf("IsOnlyDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOnlyAlphabet(t *testing.T) {

	var v string
	var got, exp bool

	v = "Bonjour"
	got, exp = IsOnlyAlphabet(v), true
	if got != exp {
		t.Errorf("IsOnlyAlphabet(%s) = %t, exp = %t", v, got, exp)
	}

	v = "Bonjour 42"
	got, exp = IsOnlyAlphabet(v), false
	if got != exp {
		t.Errorf("IsOnlyAlphabet(%s) = %t, exp = %t", v, got, exp)
	}
}

func TestFindIndexOfString(t *testing.T) {
	type args struct {
		list []string
		s    string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "standard use case",
			args: args{list: []string{"Bonjour", "je", "s'appelle", "groot"}, s: "groot"},
			want: 3,
		},
		{
			name: "not find use case",
			args: args{list: []string{"Bonjour", "je", "s'appelle", "groot"}, s: "test"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIndexOfString(tt.args.list, tt.args.s); got != tt.want {
				t.Errorf("FindIndexOfString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortString(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "standard use case",
			args: args{list: []string{"Alfred", "Momo", "Gilbert"}},
			want: []string{"Alfred", "Gilbert", "Momo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortString(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAllPosition(t *testing.T) {
	type args struct {
		s string
		r rune
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "standard use case",
			args: args{s: "This is a text and this is not", r: 'i'},
			want: []int{2, 5, 21, 24},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAllPosition(tt.args.s, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAllPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSubstringAt(t *testing.T) {
	type args struct {
		s    string
		text string
		idx  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true case",
			args: args{s: "is", text: "this is a text", idx: 2},
			want: true,
		},
		{
			name: "false case",
			args: args{s: "is", text: "this is a text", idx: 8},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubstringAt(tt.args.s, tt.args.text, tt.args.idx); got != tt.want {
				t.Errorf("IsSubstringAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
