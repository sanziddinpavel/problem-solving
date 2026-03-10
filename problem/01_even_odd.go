package problem

var arr2 = []int{1, 3, 4, 6, 55, 23, 656, 56}

func evenNumber(a []int) int {
	evenCount := 0
	for _, v := range a {
		if v%2 == 0 {

			evenCount++

		}
	}
	return evenCount
}

func oddNumber(a []int) int {
	oddCount := 0
	for _, v := range a {
		if v%2 == 1 {
			oddCount++
		}
	}
	return oddCount
}

func numbers(a []int) (int, int) {
	evenCount := 0
	oddCount := 0

	for _, v := range a {
		if v%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}

	}
	return evenCount, oddCount
}
