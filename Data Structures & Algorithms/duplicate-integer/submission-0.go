func hasDuplicate(nums []int) bool {
    res := make(map[int]bool)
    
    for _, num := range nums {
        if _, ok := res[num]; ok {
            return true
        }
        res[num] = true
    }
    return false
 }
