func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		cm := target - v
		if j, ok := m[cm]; ok {
			return []int{j, i}
		}
		m[v] = i

	}
	return []int{}
}
