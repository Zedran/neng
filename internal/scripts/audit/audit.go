// Script audit creates a standardized snapshot of all transformed words,
// which can be used to evaluate transformation quality and analyze
// the impact of new features and modifications.
//
//	go run internal/scripts/audit/audit.go
//
// Run in package's root directory.
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Zedran/neng"
	"github.com/Zedran/neng/internal/scripts/common"
)

const AUDIT_DIR string = "audit"

// Returns a string containing LF-separated forms of a single word.
// Error is emitted if any value in mods is undefined or incompatible with wc.
func buildGroup(w *neng.Word, gen *neng.Generator, wc neng.WordClass, mods neng.Mod) (string, error) {
	// Package errors are not exported. This switch ensures that the loop below
	// does not encounter any errors other than errNonComparable or errUncountable.
	switch true {
	case mods.Undefined():
		return "", errors.New("mods contain an undefined value")
	case !wc.CompatibleWith(mods):
		return "", errors.New("mods incompatible with WordClass")
	}

	var s strings.Builder

	s.WriteString(w.Word() + "\n")

	for b := neng.MOD_PLURAL; b < neng.MOD_CASE_LOWER; b <<= 1 {
		if !mods.Enabled(b) {
			continue
		}

		tw, err := gen.TransformWord(w, wc, b)

		if err != nil {
			switch wc {
			case neng.WC_ADJECTIVE, neng.WC_ADVERB:
				tw = "ncmp"
			case neng.WC_NOUN:
				tw = "unc"
			}
		}
		s.WriteString(tw + "\n")
	}
	return s.String(), nil
}

// Applies all viable grammatical transformations to every word of WordClass
// wc from the default database and writes the result to a file at audit/fname.
// Omits plural Past Simple and Present Simple forms of verbs because their
// transformation rules are trivial.
func compile(wg *sync.WaitGroup, chErr chan error, gen *neng.Generator, fname string, wc neng.WordClass) {
	const ERR_FMT string = "%s: %w"

	defer wg.Done()

	words, err := gen.All(wc)
	if err != nil {
		chErr <- fmt.Errorf(ERR_FMT, fname, err)
		return
	}

	length, err := gen.Len(wc)
	if err != nil {
		chErr <- fmt.Errorf(ERR_FMT, fname, err)
		return
	}

	mods, err := setMods(wc)
	if err != nil {
		chErr <- fmt.Errorf(ERR_FMT, fname, err)
		return
	}

	groups := make([]string, length)

	for i, w := range words {
		forms, err := buildGroup(w, gen, wc, mods)
		if err != nil {
			chErr <- fmt.Errorf(ERR_FMT, fname, err)
			return
		}
		groups[i] = forms
	}

	csum, err := common.WriteFile(filepath.Join(AUDIT_DIR, fname), groups, false)
	if err != nil {
		chErr <- fmt.Errorf(ERR_FMT, fname, err)
		return
	}

	fmt.Println(csum)
}

// Returns a Mod with every compatible grammatical bit set, depending on
// the specified WordClass. For an undefined WordClass value, returns an error.
func setMods(wc neng.WordClass) (neng.Mod, error) {
	switch wc {
	case neng.WC_ADJECTIVE, neng.WC_ADVERB:
		return neng.MOD_COMPARATIVE | neng.MOD_SUPERLATIVE, nil
	case neng.WC_NOUN:
		return neng.MOD_PLURAL, nil
	case neng.WC_VERB:
		return neng.MOD_PAST_SIMPLE | neng.MOD_PAST_PARTICIPLE | neng.MOD_PRESENT_SIMPLE | neng.MOD_GERUND, nil
	default:
		return neng.MOD_NONE, errors.New("undefined WordClass")
	}
}

func main() {
	log.SetFlags(0)

	var (
		chErr = make(chan error, 4)

		err error
		wg  sync.WaitGroup

		wordLists = map[neng.WordClass]string{
			neng.WC_ADJECTIVE: "adj",
			neng.WC_ADVERB:    "adv",
			neng.WC_NOUN:      "noun",
			neng.WC_VERB:      "verb",
		}
	)

	if err = os.Mkdir(AUDIT_DIR, 0755); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	gen, err := neng.DefaultGenerator()
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(4)

	for wc, fn := range wordLists {
		go compile(&wg, chErr, gen, fn, wc)
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
