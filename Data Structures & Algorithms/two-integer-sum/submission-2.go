func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		cm := target - v
		if _, ok := m[cm]; ok {
			return []int{m[cm], i}
		}
		m[v] = i

	}
	return []int{}
}
