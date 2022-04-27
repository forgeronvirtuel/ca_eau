package water

import (
	"reflect"
	"testing"
)

func TestGenerateCombination(t *testing.T) {
	wantsfirst := []string{
		"012",
		"013",
		"014",
		"015",
		"016",
		"017",
		"018",
		"019",
		"023",
	}
	got := GenerateCombination()

	for idx, v := range wantsfirst {
		if v != got[idx] {
			t.Errorf("idx = %d, want = %s, got = %s", idx, v, got[idx])
		}
	}

	last := got[len(got)-1]
	if last != "789" {
		t.Errorf("idx = last, want = 789, got = %s", last)
	}
}

func TestBetweenMinMax(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "min = 5, max = 10",
			args: args{min: 5, max: 10},
			want: []int{6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := BetweenMinMax(tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BetweenMinMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinimalAbsoluteDifference(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal use case 1",
			args: args{numbers: []int{5, 1, 19, 21}},
			want: 2,
		},
		{
			name: "normal use case 2",
			args: args{numbers: []int{-8, -6, 4}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimalAbsoluteDifference(tt.args.numbers); got != tt.want {
				t.Errorf("MinimalAbsoluteDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
