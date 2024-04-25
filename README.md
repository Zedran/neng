# neng, Non-Extravagant Name Generator

[![Go Reference](https://pkg.go.dev/badge/github.com/Zedran/neng.svg)](https://pkg.go.dev/github.com/Zedran/neng)
[![Go Report Card](https://goreportcard.com/badge/github.com/Zedran/neng)](https://goreportcard.com/report/github.com/Zedran/neng)

## Introduction

neng is a Golang package that can generate random names from nouns, verbs and adjectives according to user-specified pattern. It is powered by diverse collection of almost 29000 nouns, over 11000 adjectives and 6000 verbs compiled from [WordNet Lexical Database](https://wordnet.princeton.edu). Inspired by Terraria's world name generator, neng is designed to be simple yet versatile name making tool for other projects.

## Sample use

### Code

```Go
package main

import (
    "fmt"

    "github.com/Zedran/neng"
)

func main() {
    gen, _ := neng.NewGenerator()

    // <title case + noun> <Simple Present + verb> a <upper case + adjective> <upper case + noun>
    phrase, _ := gen.Phrase("%tn %Nv a %ua %un")

    // A single, transformed verb
    verb, _ := gen.Verb(neng.MOD_PAST_SIMPLE)

    // Transforming an arbitrary word
    word, _ := gen.Transform("STASH", neng.MOD_GERUND, neng.MOD_CASE_LOWER)

    fmt.Printf("Phrase -> %s\nVerb   -> %s\nWord   -> %s\n", phrase, verb, word)
}
```

### Output

```text
Phrase -> share
Verb   -> Serenade perplexes a STRAY SUPERBUG
Word   -> stashing
```

## Phrase pattern commands

**Escape character**: `%`

### Insertion

| Symbol | Description                |
|:------:|:---------------------------|
| `%`    | Inserts `%` sign           |
| `a`    | Inserts a random adjective |
| `n`    | Inserts a random noun      |
| `v`    | Inserts a random verb      |

### Transformation

Currently, no compatibility checks have been implemented. It is legal to transform any word with any modifier, it is also possible to assign more than one modifier of the same type to a word. Improper use of modifiers will therefore result in deformations.

| Symbol | Compatible with       | Package constant      | Description                |
|:------:|:---------------------:|:----------------------|:---------------------------|
| `2`    | verb                  | `MOD_PAST_SIMPLE`     | Past Simple (2nd form)     |
| `3`    | verb                  | `MOD_PAST_PARTICIPLE` | Past Participle (3rd form) |
| `N`    | verb                  | `MOD_PRESENT_SIMPLE`  | Present Simple (now)       |
| `g`    | verb                  | `MOD_GERUND`          | Gerund                     |
| `l`    | any                   | `MOD_CASE_LOWER`      | lower case                 |
| `t`    | any                   | `MOD_CASE_TITLE`      | Title Case                 |
| `u`    | any                   | `MOD_CASE_UPPER`      | UPPER CASE                 |

Symbols are used to specify transformation parameters for words within a phrase. Package constants are designed to work with "single-word" methods.

Some verbs are not correctly transformed into their past forms and gerund. Predicting when to double the final consonant has proven difficult.

## Attributions

Refer to [NOTICE.md](./NOTICE.md).

## License

This software is available under MIT License.
