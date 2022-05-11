package homework

import (
	"testing"
)

func TestAverage(t *testing.T) {
	data := []struct {
		In  [15]float32
		Out float32
	}{
		{In: [15]float32{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, Out: 2.0},
		{In: [15]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5}, Out: 4.0},
	}
	for _, d := range data {
		got := average(d.In)
		if got != d.Out {
			t.Errorf("Got %f; want %f", got, d.Out)
		}
	}
}
