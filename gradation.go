package neng

/* Returns comparative form of an adjective or an adverb (good -> better). */
func comparative(a string, adjIrr [][]string) string {
	aLine := findIrregular(a, adjIrr)
	if aLine != nil {
		return aLine[1]
	}

	return "more " + a
}

/* Returns superlative form of an adjective or an adverb (good -> best). */
func superlative(a string, adjIrr [][]string) string {
	aLine := findIrregular(a, adjIrr)
	if aLine != nil {
		return aLine[2]
	}

	return "most " + a
}
