package main

type Number interface {
	int64 | float64
}

func main() {
	intNumbers := map[string]int64{
		"first":  1,
		"second": 2,
	}

	floatNumbers := map[string]float64{
		"first":  1.1,
		"second": 2.2,
	}
	println("sum of int %v", sumNumbers(intNumbers))
	println("sum of float %v", sumNumbers(floatNumbers))
}

func sumNumbers[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}
