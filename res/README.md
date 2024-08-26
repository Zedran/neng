# Resource files

## Main resource files

These files are used to compile the embedded word lists.

| File name  | Contents                                    |
|:-----------|:--------------------------------------------|
| `adj`      | Main list of adjectives                     |
| `adj.irr`  | Adjectives with irregular graded forms      |
| `adj.ncmp` | Non-comparable adjectives                   |
| `adj.suf`  | Adjectives graded with suffixes (-er, -est) |
| `adv`      | Main list of adverbs                        |
| `adv.irr`  | Adverbs with irregular graded forms         |
| `adv.ncmp` | Non-comparable adverbs                      |
| `adv.suf`  | Adverbs graded with suffixes (-er, -est)    |
| `noun`     | Main list of nouns                          |
| `noun.irr` | Nouns with irregular plural forms           |
| `noun.plo` | Nouns that only have plural form            |
| `noun.unc` | Uncountable nouns                           |
| `verb`     | Main list of verbs                          |
| `verb.irr` | Verbs with irregular past tense forms       |

## Filters

Files in `filters` directory contain words from WordNet database that are excluded from the main resource files. Each filter is named after the main list file to which it is applied.

## Misc

| File name           | Contents                 |
|:--------------------|:-------------------------|
| `replacements.json` | List of spelling changes |
