// neng -- Non-Extravagant Name Generator
// Copyright (C) 2024  Wojciech Głąb (github.com/Zedran)
//
// This file is part of neng.
//
// neng is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 only.
//
// neng is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with neng.  If not, see <https://www.gnu.org/licenses/>.

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
	"github.com/Zedran/neng/symbols"
)

const AUDIT_DIR string = "audit"

// buildGroup returns a string containing LF-separated forms of a single word.
// Error is emitted if any value in mods is undefined or incompatible with wc.
func buildGroup(w *neng.Word, gen *neng.Generator, wc neng.WordClass, mods neng.Mod) (string, error) {
	var s strings.Builder

	tw, err := gen.TransformWord(w, wc, neng.MOD_INDEF)
	if err != nil {
		if wc != neng.WC_VERB && !errors.Is(err, symbols.ErrPluralOnly) && !errors.Is(err, symbols.ErrUncountable) {
			return "", err
		}
		tw = w.Word()
	}
	s.WriteString(tw + "\n")

	for b := neng.MOD_PLURAL; b < neng.MOD_INDEF; b <<= 1 {
		if !mods.Enabled(b) {
			continue
		}

		tw, err = gen.TransformWord(w, wc, b)

		if err != nil {
			switch wc {
			case neng.WC_ADJECTIVE, neng.WC_ADVERB:
				if !errors.Is(err, symbols.ErrNonComparable) {
					return "", err
				}
				tw = "ncmp"
			case neng.WC_NOUN:
				if !errors.Is(err, symbols.ErrUncountable) {
					return "", err
				}
				tw = "unc"
			default:
				return "", err
			}
		}
		s.WriteString(tw + "\n")
	}
	return s.String(), nil
}

// compile applies all viable grammatical transformations to every word
// of WordClass wc from the default database and writes the result to a file
// at audit/fname. Omits plural Past Simple and Present Simple forms of verbs
// because their transformation rules are trivial.
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

// setMods returns a Mod with every compatible grammatical bit set, depending on
// the specified WordClass. For an undefined WordClass value, returns an error.
// MOD_INDEF is not set by this function. buildGroup handles it in a different
// way than the rest of the Mods.
func setMods(wc neng.WordClass) (neng.Mod, error) {
	switch wc {
	case neng.WC_ADJECTIVE, neng.WC_ADVERB:
		return neng.MOD_COMPARATIVE | neng.MOD_SUPERLATIVE, nil
	case neng.WC_NOUN:
		return neng.MOD_PLURAL, nil
	case neng.WC_VERB:
		return neng.MOD_PAST_SIMPLE | neng.MOD_PAST_PARTICIPLE | neng.MOD_PRESENT_SIMPLE | neng.MOD_GERUND, nil
	default:
		return neng.MOD_NONE, symbols.ErrUndefinedWordClass
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

	gen, err := neng.DefaultGenerator(nil)
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
