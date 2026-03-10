package problem

func replaceElements(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		subArr := arr[i+1:]

		if len(subArr) == 0 {
			arr[i] = -1
			break
		}

		m := subArr[0]
		for _, v := range subArr {
			if v > m {
				m = v
			}
		}

		arr[i] = m
	}

	return arr
}
