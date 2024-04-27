package neng

/* Stores information resulting from the analysis of a given word. */
type wordInfo struct {
	// Consonant-vowel sequence of the word ('word' == 'cvcc')
	sequence string

	// Number of syllables in the word
	sylCount int
}

/* Analyzes the word and returns the outcome stored in wordInfo struct. */
func getWordInfo(word string) wordInfo {
	seq := getSequence(word)

	return wordInfo{
		sequence: seq,
		sylCount: countSyllables(word, seq),
	}
}
