package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Zedran/neng"
)

const HELP string = "Specify your pattern.\n\n" +
	"Insertions:\n" +
	"\ta - adjective         m - adverb\n" +
	"\tn - noun              v - verb\n\n" +
	"Transformations:\n" +
	"\t2 - Past Simple       3 - Past Participle    N - Present Simple\n" +
	"\tc - Comparative       g - Gerund             i - Indefinite article\n" +
	"\tp - Plural            l - lower case         s - Superlative\n" +
	"\tt - Title Case        u - UPPER CASE         _ - Indefinite (noun)\n\n" +
	"Example: %tsa %tpn that %m %Npv %in\n\n"

func main() {
	gen, err := neng.DefaultGenerator()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("%s", HELP)

	fmt.Print("pattern> ")
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			break
		}

		phrase, err := gen.Phrase(scanner.Text())
		if err != nil {
			phrase = "err: " + err.Error()
		}

		fmt.Printf("       > %s\n", phrase)
		fmt.Print("pattern> ")
	}
	fmt.Println()
}
