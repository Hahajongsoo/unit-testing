package ex1

func IsStringLong(input string) bool {
	if len(input) > 5 {
		return true
	}
	return false
	// return len(input) > 5
}

