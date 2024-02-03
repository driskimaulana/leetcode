
func isValid(s string) bool {

	bracketsPair := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	copyString := strings.Split(s, "")
	breakPoint := len(s)
	
	for i := 0; i < len(s); i++ {
		if i > breakPoint {
			breakPoint = len(s)
		}

		closeBracketIndex := -1
		nestedCount := 0
		for j := i; j < breakPoint; j++ {

			if copyString[i] == copyString[j] {
				nestedCount++
			} else if bracketsPair[copyString[i]] == copyString[j] && nestedCount != 0 {
				nestedCount--
			}

			if bracketsPair[copyString[i]] == copyString[j] {

				if nestedCount == 0 {
					copyString[i] = "0"
					copyString[j] = "0"
					breakPoint = j

					nestedCount = 0
				} else {
					if copyString[j-1] != "(" && copyString[j-1] != "{" && copyString[j-1] != "[" {
						closeBracketIndex = j
					}
					if j == breakPoint-1 && closeBracketIndex != -1 {
						copyString[i] = "0"
						copyString[closeBracketIndex] = "0"
						breakPoint = closeBracketIndex

						nestedCount = 0
					}
				}
			}

		}
	}

	for _, copyStr := range copyString {
		if copyStr != "0" {
			return false
		}
	}

	return true
}
