func isAnagram(s string, t string) bool {
	res := make(map[rune]int)

    if len(s) != len(t) {
    	return false
    }
    
    for _, char := range s {
    	res[char]++
    }
   
    for _, char := range t {
    	res[char]--
        if res[char] < 0 {
            return false
        }
    
    }
    
    return true
}