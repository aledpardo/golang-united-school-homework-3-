package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var avg float32 = 0.0
  for _, in := range input {
    avg = avg + in
  }
  avg = avg / float32(len(input))
	return avg
}

