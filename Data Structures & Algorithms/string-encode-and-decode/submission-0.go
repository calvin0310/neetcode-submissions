type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	resString := ""

	for _, str := range strs {
		resString += (strconv.Itoa(len(str)) + "#" + str)
	}

	return resString
}

func (s *Solution) Decode(encoded string) []string {
	resStrings := make([]string, 0)

	for i := 0; i < len(encoded); {
		j := i
		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])

		start := j + 1
		end := start + length
		resStrings = append(resStrings, encoded[start: end])

		i = end
	}

	return resStrings
}
