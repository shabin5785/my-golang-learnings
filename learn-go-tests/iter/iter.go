package iter

func Repeat(character string, count int) (int string) {
	var repeated string
	for i := 0; i <= count-1; i++ {
		repeated += character
	}
	return repeated
}
