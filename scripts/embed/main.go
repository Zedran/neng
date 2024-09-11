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
)

const (
	EMBED_DIR string = "embed"
	RES_DIR   string = "res"
)

// Mirrors neng.FormType
type FormType uint8

const (
	FT_REGULAR FormType = iota
	FT_IRREGULAR
	FT_PLURAL_ONLY
	FT_SUFFIXED
	FT_NON_COMPARABLE
	FT_UNCOUNTABLE
)

// Compiles the main word list and any number of supplementary lists into the embedded file
// of the same name as mainWL.
func Compile(wg *sync.WaitGroup, chErr chan error, mainFname string, supFnames ...string) {
	const errFmt = "%-5s: %w"

	defer wg.Done()

	main, err := readWL(mainFname)
	if err != nil {
		chErr <- fmt.Errorf(errFmt, mainFname, err)
		return
	}

	supplementary, err := createSupWLs(supFnames...)
	if err != nil {
		chErr <- fmt.Errorf(errFmt, mainFname, err)
		return
	}

	embed := make([]string, len(main))
	for i, w := range main {
		embed[i] = processLine(w, supplementary)
	}

	if err = writeFile(mainFname, embed); err != nil {
		chErr <- fmt.Errorf(errFmt, mainFname, err)
		return
	}

	log.Printf("%-5s: OK\n", mainFname)
}

// Cmp function for slices.BinarySearchFunc. Compares plain string and the first element
// of a comma-separated irregular line.
func cmpIrr(irr, b string) int {
	i := strings.Index(irr, ",")
	return strings.Compare(irr[:i], b)
}

// Cmp function for slices.BinarySearchFunc. Compares two strings.
func cmpStr(a, b string) int {
	return strings.Compare(a, b)
}

// Accepts the names of supplementary files of one neng.WordClass and combines their contents into a map.
func createSupWLs(fnames ...string) (map[FormType][]string, error) {
	formTypes := make(map[FormType][]string)

	for _, fn := range fnames {
		wl, err := readWL(fn)
		if err != nil {
			return nil, err
		}

		formTypes[getFormType(fn)] = wl
	}

	return formTypes, nil
}

// Determines FormType value based on file extension.
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

// Builds a single line of the embedded file.
func processLine(word string, supWLs map[FormType][]string) string {
	for ft, wl := range supWLs {
		var cmp func(string, string) int

		if ft == FT_IRREGULAR {
			cmp = cmpIrr
		} else {
			cmp = cmpStr
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

// Reads word list from RES_DIR and returns a slice of words.
func readWL(fname string) ([]string, error) {
	stream, err := os.ReadFile(filepath.Join(RES_DIR, fname))
	if err != nil {
		return nil, err
	}

	return strings.Split(string(stream), "\n"), nil
}

// Writes lines to the specified embedded file.
func writeFile(fname string, lines []string) error {
	return os.WriteFile(filepath.Join(EMBED_DIR, fname), []byte(strings.Join(lines, "\n")), 0644)
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
		go Compile(&wg, chErr, main, sup...)
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
