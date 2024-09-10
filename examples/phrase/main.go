package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Zedran/neng"
)

const HELP string = "Specify your pattern.\n\n" +
	"a - adjective         m - adverb\n" +
	"n - noun              v - verb\n\n" +
	"2 - Past Simple       3 - Past Participle    N - Present Simple\n" +
	"c - Comparative       g - Gerund             p - Plural\n" +
	"l - lower case        s - Superlative        t - Title Case\n" +
	"u - UPPER CASE\n\n" +
	"Example: %tsa %tpn that %m %Npv the %n\n\n"

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
