package utils

func GetAverageArrayInt(array []int) float32 {
	n := len(array)
	sum := 0

	for i := 0; i < n; i++ {
		sum += (array[i])
	}

	return (float32(sum) / float32(n))
}

func GetAverageArrayFloat(array []float32) float32 {
	n := len(array)
	var sum float32 = 0

	for i := 0; i < n; i++ {
		sum += (array[i])
	}

	return (sum / float32(n))
}

func GetMinMaxArray(array []int) (int, int) {
	min := 1000
	max := 0
	for i := 0; i < len(array); i++ {
		if max < array[i] {
			max = array[i]
		} else if min > array[i] {
			min = array[i]
		}
	}

	return min, max
}
