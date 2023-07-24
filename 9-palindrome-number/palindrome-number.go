func isPalindrome(x int) bool {
  xString := strconv.Itoa(x)
	xRunes := []rune(xString)
	j := len(xRunes) - 1
	for i := 0; i < len(xRunes); i++ {
		if xRunes[i] != xRunes[j] {
			break
		}
		j--
	}

	return j == -1
}