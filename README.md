# neng, Non-Extravagant Name Generator

[![Go Reference](https://pkg.go.dev/badge/github.com/Zedran/neng.svg)](https://pkg.go.dev/github.com/Zedran/neng)
[![Go Report Card](https://goreportcard.com/badge/github.com/Zedran/neng)](https://goreportcard.com/report/github.com/Zedran/neng)

## Introduction

neng is a Golang package that can generate random English phrases from nouns, verbs, adverbs and adjectives according to user-specified pattern. It is powered by diverse collection of 41000 words compiled from [WordNet Lexical Database](https://wordnet.princeton.edu). Inspired by Terraria's world name generator, neng is designed to be simple yet versatile name making tool for other projects.

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
    gen, _ := neng.DefaultGenerator(nil)

    // <title case + noun> <Present Simple + verb> the <upper + adjective> <upper + noun>.
    phrase, _ := gen.Phrase("%tn %Nv the %ua %un.")

    // A single, transformed verb
    verb, _ := gen.Verb(neng.MOD_PAST_SIMPLE)

    // Transforming an arbitrary word
    word, _ := gen.Transform("involve", neng.WC_VERB, neng.MOD_GERUND|neng.MOD_CASE_TITLE)

    fmt.Printf("Phrase -> %s\nVerb   -> %s\nWord   -> %s\n", phrase, verb, word)
}
```

### Output

```text
Phrase -> Serenade perplexes the STRAY SUPERBUG.
Verb   -> shared
Word   -> Involving
```

## Phrase pattern commands

neng's phrase generation syntax resembles C-style string format specifiers. The user provides a pattern consisting of commands, preceded by the escape character, to insert (and optionally transform) randomly generated words. A regular text can be mixed with command syntax.

**Escape character**: `%`

### Insertion

| Symbol | WordClass      | Description                |
|:------:|:--------------:|:---------------------------|
| `%`    |                | Inserts `%` sign           |
| `a`    | `WC_ADJECTIVE` | Inserts a random adjective |
| `m`    | `WC_ADVERB`    | Inserts a random adverb    |
| `n`    | `WC_NOUN`      | Inserts a random noun      |
| `v`    | `WC_VERB`      | Inserts a random verb      |

`WordClass` values are required by some of the Generator's methods to recognize parts of speech.

### Transformation

Transformations can only be applied to compatible parts of speech.

Symbols are used to request transformations for words within a phrase. Constants of type [`Mod`](./mod.go#L21) are designed to work with "single-word" methods.

| Symbol | Compatible with          | Mod                   | Description                  |
|:------:|:------------------------:|:----------------------|:-----------------------------|
| `2`    | verb                     | `MOD_PAST_SIMPLE`     | Past Simple (2nd form)       |
| `3`    | verb                     | `MOD_PAST_PARTICIPLE` | Past Participle (3rd form)   |
| `N`    | verb                     | `MOD_PRESENT_SIMPLE`  | Present Simple (now)         |
| `c`    | adjective, adverb        | `MOD_COMPARATIVE`     | Comparative (better)         |
| `f`    | any                      | `MOD_CASE_SENTENCE`   | Sentence case (first letter) |
| `g`    | verb                     | `MOD_GERUND`          | Gerund                       |
| `i`    | adjective, adverb, noun* | `MOD_INDEF`           | Indefinite adjective (a, an) |
| `_`    | noun                     | `MOD_INDEF_SILENT`    | Silent indefinite**          |
| `l`    | any                      | `MOD_CASE_LOWER`      | lower case                   |
| `o`    | noun                     | `MOD_POSSESSIVE`      | Possessive form (owner)      |
| `p`    | noun, verb***            | `MOD_PLURAL`          | Plural form                  |
| `s`    | adjective, adverb        | `MOD_SUPERLATIVE`     | Superlative (best)           |
| `t`    | any                      | `MOD_CASE_TITLE`      | Title Case                   |
| `u`    | any                      | `MOD_CASE_UPPER`      | UPPER CASE                   |

\* `MOD_INDEF` is not compatible with `MOD_PLURAL` and `MOD_SUPERLATIVE`.

\*\* `MOD_INDEF_SILENT` ensures that the noun is grammatically compatible with an indefinite article (not uncountable, not plural-only), but does not modify it in any way. It is useful in phrase patterns such as `%ia %_n`, where the indefinite article belongs to the noun, but it stands before the adjective describing the noun. If `Generator.TransformWord` method receives silent indefinite, it does nothing to the provided word, but it still returns an error in case of incompatibility.

\*\*\* `MOD_PLURAL` is only compatible with verbs when combined with `MOD_PAST_SIMPLE` or `MOD_PRESENT_SIMPLE`.

`Mod` values form two conceptual categories: grammar modifiers and case modifiers. Only one modifier from each category may be applied to any given word. If multiple modifiers of the same kind are specified, the one with the lowest value is applied. The above-mentioned verb transformations with `MOD_PLURAL` and `MOD_INDEF` are exceptions to this rule.

## State of the vocabulary

Generator's default vocabulary consists of:

* 23000 nouns
* 10000 adjectives
* 6000 verbs
* 2000 adverbs

Original WordNet lists have been thoroughly vetted. I have strived to remove any words that are offensive, too specific (chemistry, medicine) or relate to topics that are considered sensitive, controversial or fear-inducing. However, I am not native English speaker and the database is quite large, so it is likely I have missed something. If you find any unsuitable words, I will be happy to hear from you.

If the embedded database does not meet your requirements, you can provide neng with your own word lists. To ensure the accuracy of transformations, I recommend that your custom vocabulary remains a subset of the embedded one.

## Attributions

Refer to [NOTICE.md](./NOTICE.md).

## License

This project is available under GPL-3.0 License.

Originally, neng was released under MIT License. However, during development, I included materials licensed under CC BY-SA 4.0, unaware that its share-alike clause is incompatible with MIT License. In order to satisfy "the newfound" requirements, I relicensed neng under GPL-3.0, starting with v0.16.0. Versions with licensing conflicts (v0.8.2 through v0.15.4) have been retracted from Go package index.
