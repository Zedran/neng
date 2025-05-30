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

// Script embed builds the embedded files from resource files
// created with [github.com/Zedran/neng/internal/scripts/res].
//
//	go run internal/scripts/embed/embed.go
//
// Run in package's root directory.
package main

import (
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
	EMBED_DIR string = "embed"
	RES_DIR   string = "res"
)

// FormType mirrors neng.FormType. The original cannot be used, because
// importing neng causes compile-time error if the embedded files are missing.
type FormType uint8

const (
	FT_REGULAR FormType = iota
	FT_IRREGULAR
	FT_PLURAL_ONLY
	FT_SUFFIXED
	FT_NON_COMPARABLE
	FT_UNCOUNTABLE
)

// cmpIrr is a cmp function for slices.BinarySearchFunc. Compares plain string
// and the first element of a comma-separated irregular line.
func cmpIrr(irr, b string) int {
	i := strings.Index(irr, ",")
	return strings.Compare(irr[:i], b)
}

// compile builds the main word list and any number of supplementary lists into
// the embedded file stored in EMBED_DIR/mainFname.
func compile(wg *sync.WaitGroup, chErr chan error, mainFname string, supFnames ...string) {
	const ERR_FMT = "%s: %w"

	defer wg.Done()

	main, err := common.ReadFile(filepath.Join(RES_DIR, mainFname))
	if err != nil {
		chErr <- fmt.Errorf(ERR_FMT, mainFname, err)
		return
	}

	supplementary, err := readSupWLs(supFnames...)
	if err != nil {
		chErr <- fmt.Errorf(ERR_FMT, mainFname, err)
		return
	}

	embed := make([]string, len(main))
	for i, w := range main {
		embed[i] = processLine(w, supplementary)
	}

	csum, err := common.WriteFile(filepath.Join(EMBED_DIR, mainFname), embed, false)
	if err != nil {
		chErr <- fmt.Errorf(ERR_FMT, mainFname, err)
		return
	}

	fmt.Println(csum)
}

// getFormType determines FormType value based on file extension.
func getFormType(fname string) FormType {
	switch filepath.Ext(fname) {
	case ".irr":
		return FT_IRREGULAR
	case ".plo":
		return FT_PLURAL_ONLY
	case ".suf":
		return FT_SUFFIXED
	case ".ncmp":
		return FT_NON_COMPARABLE
	case ".unc":
		return FT_UNCOUNTABLE
	default:
		return FT_REGULAR
	}
}

// processLine builds a single line of the embedded file.
func processLine(word string, supWLs map[FormType][]string) string {
	for ft, wl := range supWLs {
		var cmp func(string, string) int

		if ft == FT_IRREGULAR {
			cmp = cmpIrr
		} else {
			cmp = strings.Compare
		}

		i, found := slices.BinarySearchFunc(wl, word, cmp)
		if !found {
			continue
		}

		var irr string = ""
		if ft == FT_IRREGULAR {
			irr, _ = strings.CutPrefix(wl[i], word) // == ",irr1[,irr2]"
		}

		return fmt.Sprintf("%d%s%s", ft, word, irr)
	}

	return "0" + word
}

// readSupWLs accepts the file names of supplementary files of a single
// main list and combines their contents into a map.
func readSupWLs(fnames ...string) (map[FormType][]string, error) {
	sup := make(map[FormType][]string)

	for _, fn := range fnames {
		wl, err := common.ReadFile(filepath.Join(RES_DIR, fn))
		if err != nil {
			return nil, err
		}

		sup[getFormType(fn)] = wl
	}

	return sup, nil
}

func main() {
	log.SetFlags(0)

	if err := os.Mkdir(EMBED_DIR, 0755); err != nil && !os.IsExist(err) {
		log.Fatalf("Could not create embed directory: %v", err)
	}

	var (
		chErr = make(chan error, 4)

		err error
		wg  sync.WaitGroup

		res = map[string][]string{
			"adj":  {"adj.irr", "adj.ncmp", "adj.suf"},
			"adv":  {"adv.irr", "adv.ncmp", "adv.suf"},
			"noun": {"noun.irr", "noun.plo", "noun.unc"},
			"verb": {"verb.irr"},
		}
	)

	wg.Add(4)

	for main, sup := range res {
		go compile(&wg, chErr, main, sup...)
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
