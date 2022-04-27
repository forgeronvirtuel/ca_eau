package water

import (
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
