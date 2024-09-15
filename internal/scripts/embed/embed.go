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

	"github.com/Zedran/neng/internal/common"
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

// Cmp function for slices.BinarySearchFunc. Compares plain string and the first element
// of a comma-separated irregular line.
func cmpIrr(irr, b string) int {
	i := strings.Index(irr, ",")
	return strings.Compare(irr[:i], b)
}

// Compiles the main word list and any number of supplementary lists into the embedded file
// stored in EMBED_DIR/mainFname.
func compile(wg *sync.WaitGroup, chErr chan error, mainFname string, supFnames ...string) {
	const ERR_FMT = "%-5s: %w"

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

	if err = common.WriteFile(filepath.Join(EMBED_DIR, mainFname), embed, false); err != nil {
		chErr <- fmt.Errorf(ERR_FMT, mainFname, err)
		return
	}

	log.Printf("%-5s: OK\n", mainFname)
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

// Accepts the file names of supplementary files of a single main list and combines their contents into a map.
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
