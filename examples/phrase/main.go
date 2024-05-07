package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Zedran/neng"
)

func main() {
	gen, err := neng.DefaultGenerator()
	if err != nil {
		log.Fatal(err)
	}

	var (
		pattern string
		phrase  string
		scanner = bufio.NewScanner(os.Stdin)

		help string = "Specify your phrase.\n\n" +
			"a - adjective         n - noun               v - verb\n\n" +
			"2 - Past Simple       3 - Past Participle    N - Present Simple\n" +
			"g - Gerund            p - Plural             l - lower case\n" +
			"t - Title Case        u - UPPER CASE\n\n" +
			"Example: %ta %tpn that %Npv the %n\n\n"
	)

	fmt.Print(help)

	for {
		fmt.Print("phrase> ")

		if scanner.Scan() {
			pattern = scanner.Text()
		}

		phrase, err = gen.Phrase(pattern)
		if err != nil {
			phrase = "err: " + err.Error()
		}

		fmt.Printf("      > %s\n", phrase)
	}
}
