// getRandomCharacter returns a random character from the given string.
func getRandomCharacter(input string) string {
	generator := getSeed()
	index := generator.Intn(len(input))
	return string(input[index])
}