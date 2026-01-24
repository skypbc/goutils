package tests

var testData = []int{1, 2, 3, 4, 5}

func getTestData() []int {
	res := make([]int, len(testData))
	copy(res, testData)
	return res
}
