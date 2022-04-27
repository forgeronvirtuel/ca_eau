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
