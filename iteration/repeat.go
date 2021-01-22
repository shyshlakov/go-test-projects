package iteration

// Repeat takes character and returns string
func Repeat(character string, count int) string {
	var repeated string

	if count == 0 {
		count = 5
	}
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
