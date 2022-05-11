package homework

import "testing"

func TestReverse(t *testing.T) {
	data := []struct {
		In  []int64
		Out []int64
	}{
		{In: []int64{1, 2, 3, 4, 5}, Out: []int64{5, 4, 3, 2, 1}},
	}
	for _, d := range data {
		got := reverse(d.In)
		if !Equal(got, d.Out) {
			t.Errorf("Got %p; want %p", got, d.Out)
		}
	}
}

func Equal(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
