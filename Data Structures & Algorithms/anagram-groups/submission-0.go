import (
    "slices"
)


func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
    m := make(map[string][]string)

    if len(strs) == 1  {
   	    res = append(res, strs)
        return res
    }
    
    for _, word := range strs {
		runeWord := []rune(word)
		
	   	slices.Sort(runeWord)
		newWord := string(runeWord)
		m[newWord] = append(m[newWord], word)
        
    }
    
    for _, words := range m {
    	res = append(res, words)
    }
    
    return res
    
}