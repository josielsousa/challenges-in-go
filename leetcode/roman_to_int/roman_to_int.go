package main

func romanToInt(s string) int {
	romans := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	sufix := map[string]string{
		"I": "VX", "X": "LC", "C": "DM",
	}

	var sum int

	for i := len(s) - 1; i >= 0; i-- {
		key := string(s[i])

		if svdataset, ok := sufix[key]; ok && i+1 < len(s) {
			if svdataset[0] == s[i+1] || svdataset[1] == s[i+1] {
				sum -= romans[key]
				continue
			}
		}

		sum += romans[key]
	}

	return sum
}
