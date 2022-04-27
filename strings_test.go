package water

import "testing"

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
