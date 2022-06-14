package iteration

func Repeat(character string, times int) string {
	// Up until now we have been using :=, which is shorthand for both declaring and initializing variables.
	// Here, we are declaring it only, so we don't need to initialize anything.
	var result string
	for i := 0; i < times; i++ {
		result += character
	}
	return result
}
