func groupAnagrams(strs []string) [][]string {
	buket := make(map[[26]int][]string)

	for _, s := range strs{
		var key [26]int

		for _, c := range s{
			key[c-'a']++
		}

		buket[key] = append(buket[key], s)
	}

	result := [][]string{}

	for _, b := range buket{
		result = append(result, b)
	}
	return result
}
