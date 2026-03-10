package problem

import "fmt"

var arr = []int{1, 3, 4, 6, 55, 23, 656}

func max(a []int) (max int, max2 int, max3 int) {
	if len(a) == 0 {
		return
	}

	max = a[0]

	for _, v := range a {
		if v > max {
			max = v
		}

	}
	max2 = a[0]
	for _, v := range a {
		if v < max && v > max2 {
			max2 = v

		}
	}
	max3 = a[0]
	for _, v := range a {
		if v < max && v < max2 && v > max3 {
			max3 = v
		}
	}

	return
}

func Min(a []int) int {
	if len(a) == 0 {
		return 0
	}

	min := a[0]

	for _, v := range a {
		if v < min {
			min = v
		}

	}

	return min
}
func MaxAndMinValue() {
	fmt.Println(max(arr))
	fmt.Println(Min(arr))
}
