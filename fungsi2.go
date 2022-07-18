package main

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var average = float64(total) / float64(len(numbers))
	return average
}
