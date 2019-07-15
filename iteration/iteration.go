package iteration

func Repeat(character string, repeatCount int) string {
	// const repeatCount = 5
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
