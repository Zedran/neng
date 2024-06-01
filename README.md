# neng, Non-Extravagant Name Generator

[![Go Reference](https://pkg.go.dev/badge/github.com/Zedran/neng.svg)](https://pkg.go.dev/github.com/Zedran/neng)
[![Go Report Card](https://goreportcard.com/badge/github.com/Zedran/neng)](https://goreportcard.com/report/github.com/Zedran/neng)

## Introduction

neng is a Golang package that can generate random English phrases from nouns, verbs, adverbs and adjectives according to user-specified pattern. It is powered by diverse collection of 43000 words compiled from [WordNet Lexical Database](https://wordnet.princeton.edu). Inspired by Terraria's world name generator, neng is designed to be simple yet versatile name making tool for other projects.

If the embedded word database does not meet your requirements, you can provide neng with your own word lists.

## Showcase

```text
go run github.com/Zedran/neng/examples/phrase@latest
```

## Sample use

### Code

```Go
package main

import (
    "fmt"

    "github.com/Zedran/neng"
)

func main() {
    gen, _ := neng.DefaultGenerator()

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
Phrase -> Serenade perplexes a STRAY SUPERBUG
Verb   -> share
Word   -> stashing
```

## Phrase pattern commands

**Escape character**: `%`

### Insertion

| Symbol | Description                |
|:------:|:---------------------------|
| `%`    | Inserts `%` sign           |
| `a`    | Inserts a random adjective |
| `m`    | Inserts a random adverb    |
| `n`    | Inserts a random noun      |
| `v`    | Inserts a random verb      |

### Transformation

Currently, no compatibility checks have been implemented. It is legal to transform any word with any modifier, it is also possible to assign more than one modifier of the same type to a word. Improper use of modifiers will therefore result in deformations.

| Symbol | Compatible with       | Package constant      | Description                |
|:------:|:---------------------:|:----------------------|:---------------------------|
| `2`    | verb                  | `MOD_PAST_SIMPLE`     | Past Simple (2nd form)     |
| `3`    | verb                  | `MOD_PAST_PARTICIPLE` | Past Participle (3rd form) |
| `N`    | verb                  | `MOD_PRESENT_SIMPLE`  | Present Simple (now)       |
| `c`    | adjective, adverb     | `MOD_COMPARATIVE`     | Comparative (better)       |
| `g`    | verb                  | `MOD_GERUND`          | Gerund                     |
| `l`    | any                   | `MOD_CASE_LOWER`      | lower case                 |
| `p`    | noun, verb            | `MOD_PLURAL`          | Plural form                |
| `s`    | adjective, adverb     | `MOD_SUPERLATIVE`     | Superlative (best)         |
| `t`    | any                   | `MOD_CASE_TITLE`      | Title Case                 |
| `u`    | any                   | `MOD_CASE_UPPER`      | UPPER CASE                 |

Symbols are used to specify transformation parameters for words within a phrase. Package constants are designed to work with "single-word" methods.

## State of the vocabulary

Generator's default vocabulary consists of:

* 25000 nouns
* 10000 adjectives
* 6000 verbs
* 2000 adverbs

Original WordNet lists have been thoroughly vetted. I have strived to remove any words that are offensive, too specific (chemistry, medicine) or relate to topics that are considered sensitive, controversial or fear-inducing. However, I am not native English speaker and the database is quite large, so it is likely I have missed something. If you find any unsuitable words, I will be happy to hear from you.

## Attributions

Refer to [NOTICE.md](./NOTICE.md).

## License

This software is available under MIT License.
