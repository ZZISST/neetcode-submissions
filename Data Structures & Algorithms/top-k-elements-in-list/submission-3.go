import (
    "slices"
)

func topKFrequent(nums []int, k int) []int {
    res := []int{}

    m := make(map[int]int)
    mres := make(map[int][]int)
    mkeys := []int{}
    var maxIn int

    for _, num := range nums {
        m[num]++
        maxIn = max(maxIn, m[num])
    }
    
    for key, val := range m {
     	mkeys = append(mkeys, val)
    	mres[val] = append(mres[val], key)
    }
    
    slices.Sort(mkeys)
    slices.Reverse(mkeys)
    
    
    for _, key := range mkeys  {
    	if k > len(res) {
     		if  _, ok := mres[key]; ok{
			   	if len(mres[key]) > 1 {
					num := mres[key][0]
					slices.Delete(mres[key], 0, 1)
			  		res = append(res, num)
			   	}
				if len(mres[key]) == 1 {
					num := mres[key][0]
			  		res = append(res, num)
				}
       		}
     	} else {
            break
      	}
      
    }
    
    slices.Sort(res)

    return res
}
