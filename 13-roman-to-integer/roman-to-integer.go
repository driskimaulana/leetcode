func romanToInt(s string) int {
  result := 0
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "I":
			if i != len(s)-1 {
				if string(s[i+1]) == "V" {
					result += 4
					i++
					break
				}
				if string(s[i+1]) == "X" {
					result += 9
					i++
					break
				}
			}
			result += 1
			break
		case "V":
			result += 5
			break
		case "X":
			if i != len(s)-1 {
				if string(s[i+1]) == "L" {
					result += 40
					i++
					break
				}
				if string(s[i+1]) == "C" {
					result += 90
					i++
					break
				}
			}
			result += 10
			break
		case "L":
			result += 50
			break
		case "C":
			if i != len(s)-1 {
				if string(s[i+1]) == "D" {
					result += 400
					i++
					break
				}
				if string(s[i+1]) == "M" {
					result += 900
					i++
					break
				}
			}
			result += 100
			break
		case "D":
			result += 500
			break
		case "M":
			result += 1000
			break
		}
	}

	return result
}