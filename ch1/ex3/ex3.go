package ex3

var WasLastStringLong bool

func IsStringLong(input string) bool {
	result := len(input) > 5
	WasLastStringLong = result
	return result
}
