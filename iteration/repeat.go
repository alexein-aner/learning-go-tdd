package iteration

func Repeat(char string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated = repeated + char
	}
	return repeated
}
