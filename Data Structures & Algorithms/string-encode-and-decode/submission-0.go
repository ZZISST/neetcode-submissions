import "unicode/utf8"

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encoded_string string

	for i := 0; i < len(strs); i++ {
		c := utf8.RuneCountInString(strs[i])
		encoded_string += fmt.Sprintf("%d#%s", c, strs[i])
	}
	return encoded_string
}

func (s *Solution) Decode(encoded string) []string {
	var strs []string
	var lenW string
	for i := 0; i < utf8.RuneCountInString(encoded); {
		if encoded[i] != '#' {
			lenW += string(encoded[i])
			i++
		} else {
			i++
			lenWint, err := strconv.Atoi(lenW)
			if err == nil {
				var word string
				for j := range lenWint {
					word += string(encoded[i+j])
				}
				strs = append(strs, word)
				lenW = ""
				i += lenWint
			}
		}
	}
	return strs
}

