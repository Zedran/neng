# neng, Non-Extravagant Name Generator

## Introduction

neng is a Golang package that can generate random names from nouns and adjectives according to user-specified pattern. It is powered by a diverse collection of almost 29000 nouns and over 11000 adjectives compiled from [WordNet Lexical Database](https://wordnet.princeton.edu). Inspired by Terraria's world name generator, neng is designed to be simple yet versatile name making tool for other projects.

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

    // A single adjective
    fmt.Println(gen.Adjective())

    // A single noun
    fmt.Println(gen.Noun())

    // <adjective> <noun> of the <noun>
    fmt.Println(gen.Phrase("%a %n of the %n"))
}
```

### Output

```text
bituminous
carnosaur
revolutionary conversation of the bacon
```

## Phrase pattern commands

Escape character: `%`

| Command | Description          |
|:-------:|:---------------------|
| `%%`    | Inserts `%` sign     |
| `%a`    | Inserts an adjective |
| `%n`    | Inserts a noun       |

## Attributions

Refer to [NOTICE.md](./NOTICE.md).

## License

This software is available under MIT License.
