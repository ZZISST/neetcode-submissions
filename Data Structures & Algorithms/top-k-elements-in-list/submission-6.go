func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)
	highK := make(map[int][]int)
	counts := []int{}
	res := []int{}

	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		highK[v] = append(highK[v], k)
	}

	for k, _ := range highK {
		counts = append(counts, k)
	}

	fmt.Println(m, highK, counts)

	sort.Slice(counts, func(i, j int) bool { return counts[i] > counts[j] })

	fmt.Println(counts)

	i := 0
	for len(res) < k {
		for _, v := range highK[counts[i]]{
		if len(res) == k {
			break
		}
		res = append(res, v)
		}
		i++
	}

	return res
}