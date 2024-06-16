// getSeed initializes a new random number generator with a seed based on the current time
func getSeed() *rand.Rand {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	generator := rand.New(source)
	return generator
}