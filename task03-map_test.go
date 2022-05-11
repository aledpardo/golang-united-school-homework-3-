package homework

import "testing"

func TestSortMapValues(t *testing.T) {
	var equals func(a, b []string) bool
	equals = func(a, b []string) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	data := []struct {
		In  map[int]string
		Out []string
	}{
		{In: map[int]string{2: "a", 0: "b"}, Out: []string{"b", "a"}},
		// { In: map[int]string {2: "a", 0: "b", 1: "c"}, Out: ["b", "c", "a"]},
	}
	for _, d := range data {
		got := sortMapValues(d.In)
		if !equals(got, d.Out) {
			t.Errorf("wrong sorting")
		}
	}
}
