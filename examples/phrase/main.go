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
	"\to - Possessive form   p - Plural             s - Superlative\n" +
	"\tl - lower case        f - Sentence case      t - Title Case\n" +
	"\tu - UPPER CASE        _ - Indefinite (noun)\n\n" +
	"Example: %tsa %tpn that %m %Npv %in\n\n"

func main() {
	gen, err := neng.DefaultGenerator(nil)
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
