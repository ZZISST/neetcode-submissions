func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)
	var result = [][]string{}
	for _, s := range strs {
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		realS := string(b)

		res[realS] = append(res[realS], s)
	}

	for _, v := range res {
		result = append(result, v)
	}

	return result

}
