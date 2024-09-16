// This script builds the resource files from WordNet source.
//
//	go run internal/scripts/res/res.go
//
// Run in package's root directory.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	"github.com/Zedran/neng/internal/scripts/common"
)

const (
	RES_DIR     string = "res"
	FILTERS_DIR string = "res/filters"
	REPL_FILE   string = "res/misc/replacements.json"
	WNET_DIR    string = "res/wordnet"
)

// Appends irregular verbs that are missing from lines.
func appendMissingIrr(lines []string, irr []string) []string {
	for _, iw := range irr {
		i := strings.Index(iw, ",")
		if _, found := slices.BinarySearch(lines, iw[:i]); !found {
			lines = append(lines, iw[:i])
		}
	}
	return lines
}

// Removes lines that are present in filter.
func applyFilter(lines, filter []string) []string {
	for _, f := range filter {
		if i, found := slices.BinarySearch(lines, f); found {
			lines[i] = ""
		}
	}

	slices.Sort(lines)
	return slices.Compact(lines)[1:]
}

// Compiles the main resource list from WordNet source file. Accepts the following arguments:
//   - srcFname     - name of the source file, without "data." prefix
//   - replacements - word replacements for a given srcFname (res/misc/replacements.json)
func compile(wg *sync.WaitGroup, chErr chan error, srcFname string, replacements map[string]string) {
	const (
		LICENSE_OFFSET int    = 29
		ERR_FMT        string = "%s: %w"
	)

	defer wg.Done()

	lines, err := common.ReadFile(filepath.Join(WNET_DIR, "data."+srcFname))
	if err != nil {
		chErr <- fmt.Errorf(ERR_FMT, srcFname, err)
		return
	}

	filter, err := common.ReadFile(filepath.Join(FILTERS_DIR, srcFname+".filter"))
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			chErr <- fmt.Errorf(ERR_FMT, srcFname, err)
			return
		}
		log.Printf("%s: %s.filter file not found. Proceeding without it.\n", srcFname, srcFname)
	}

	// Strip license
	lines = lines[LICENSE_OFFSET:]

	discardMetadata(lines)

	for i, ln := range lines {
		if len(ln) == 1 || containsChars(ln) || isProperNoun(ln) {
			lines[i] = ""
			continue
		}
		stripParentheses(&lines[i])
	}

	if srcFname == "verb" {
		irr, err := common.ReadFile(filepath.Join(RES_DIR, srcFname+".irr"))
		if err != nil {
			chErr <- fmt.Errorf(ERR_FMT, srcFname, err)
			return
		}
		lines = appendMissingIrr(lines, irr)
	}

	// Remove duplicates
	slices.Sort(lines)
	lines = slices.Compact(lines)[1:] // remove empty line at index 0

	lines = applyFilter(lines, filter)
	replaceEntries(lines, replacements)

	csum, err := common.WriteFile(filepath.Join(RES_DIR, srcFname), lines, true)
	if err != nil {
		chErr <- fmt.Errorf(ERR_FMT, srcFname, err)
		return
	}

	fmt.Println(csum)
}

// Returns true if line contains any of the following types of entries:
//   - words containing apostrophes
//   - compound words (separated with '-')
//   - words containing numbers
//   - multi-word entries (separated with '_')
func containsChars(line string) bool {
	return strings.ContainsAny(line, "'-0123456789_")
}

// Extract the words from the surrounding metadata.
func discardMetadata(lines []string) {
	const WORD_COL int = 4

	for i := range lines {
		s := strings.Split(lines[i], " ")

		if len(s) >= WORD_COL {
			lines[i] = s[WORD_COL]
		}
	}
}

// Returns true if the line contains a proper noun or an adjective derived from a proper noun.
func isProperNoun(line string) bool {
	for _, c := range line {
		// Every letter must be checked - consider deVries
		if c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}

// Reads JSON file into a container v.
func readJSON(fname string, v any) error {
	stream, err := os.ReadFile(fname)
	if err != nil {
		return err
	}
	return json.Unmarshal(stream, v)
}

// Modifies specific words in lines.
func replaceEntries(lines []string, replacements map[string]string) {
	for old, new := range replacements {
		if i, found := slices.BinarySearch(lines, old); found {
			lines[i] = new
		}
	}
}

// Removes parenthesized content from the end of the string.
func stripParentheses(s *string) {
	i := strings.Index(*s, "(")
	if i > -1 {
		*s = (*s)[:i]
	}
}

func main() {
	log.SetFlags(0)

	var (
		chErr = make(chan error, 4)

		wg sync.WaitGroup

		src          = []string{"adj", "adv", "noun", "verb"}
		replacements = make(map[string]map[string]string)
	)

	err := readJSON(REPL_FILE, &replacements)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatal(err)
		}
		err = nil
		log.Printf("%s not found. Proceeding without it.\n", REPL_FILE)
	}

	wg.Add(4)

	for _, fname := range src {
		go compile(&wg, chErr, fname, replacements[fname])
	}

	wg.Wait()
	close(chErr)

	for e := range chErr {
		err = errors.Join(err, e)
	}

	if err != nil {
		log.Fatal(err)
	}
}
