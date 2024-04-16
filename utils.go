package neng

import (
	"math/rand"
	"strings"
)

/* Loads a word list from path. Returns error if the file is not found. */
func loadWords(path string) ([]string, error) {
	stream, err := efs.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(stream), "\n"), nil
}

/* Returns a random item from s. */
func randItem(s []string) string {
	return s[rand.Intn(len(s))]
}
