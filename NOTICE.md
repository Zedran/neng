# Notices and information

This software incorporates materials licensed by third parties.

## Wiktionary

### Source

* Wiktionary. Category: English countable nouns. Available at https://en.wiktionary.org/wiki/Category:English_countable_nouns \[2024-06-03] Wikimedia Foundation. 2022.
* Wiktionary. Category: English uncomparable adjectives. Available at https://en.wiktionary.org/wiki/Category:English_uncomparable_adjectives \[2024-07-25] Wikimedia Foundation. 2022.
* Wiktionary. Category: English uncomparable adverbs. Available at https://en.wiktionary.org/wiki/Category:English_uncomparable_adverbs \[2024-07-25] Wikimedia Foundation. 2022.
* Wiktionary. Category: English uncountable nouns. Available at https://en.wiktionary.org/wiki/Category:English_uncountable_nouns \[2024-06-03] Wikimedia Foundation. 2022.

### License

[Creative Commons Attribution-ShareAlike License](https://creativecommons.org/licenses/by-sa/4.0)

### Used material

* Page titles in the "English countable nouns" and "English uncountable nouns" categories were used to compile `res/noun.unc`.
* Page titles in the "English uncomparable adjectives" category were used to compile `res/adj.ncmp`.
* Page titles in the "English uncomparable adverbs" category were used to compile `res/adv.ncmp`.

## WordNet

### Source

WordNet: A Lexical Database for English v3.1. Available at https://wordnet.princeton.edu/download/current-version \[2024-04-15] Princeton University. 2011.

### License

```text
WordNet Release 3.1

This software and database is being provided to you, the LICENSEE, by  
Princeton University under the following license.  By obtaining, using  
and/or copying this software and database, you agree that you have  
read, understood, and will comply with these terms and conditions.:  
  
Permission to use, copy, modify and distribute this software and  
database and its documentation for any purpose and without fee or  
royalty is hereby granted, provided that you agree to comply with  
the following copyright notice and statements, including the disclaimer,  
and that the same appear on ALL copies of the software, database and  
documentation, including modifications that you make for internal  
use or for distribution.  
  
WordNet 3.1 Copyright 2011 by Princeton University.  All rights reserved.  
  
THIS SOFTWARE AND DATABASE IS PROVIDED "AS IS" AND PRINCETON  
UNIVERSITY MAKES NO REPRESENTATIONS OR WARRANTIES, EXPRESS OR  
IMPLIED.  BY WAY OF EXAMPLE, BUT NOT LIMITATION, PRINCETON  
UNIVERSITY MAKES NO REPRESENTATIONS OR WARRANTIES OF MERCHANT-  
ABILITY OR FITNESS FOR ANY PARTICULAR PURPOSE OR THAT THE USE  
OF THE LICENSED SOFTWARE, DATABASE OR DOCUMENTATION WILL NOT  
INFRINGE ANY THIRD PARTY PATENTS, COPYRIGHTS, TRADEMARKS OR  
OTHER RIGHTS.  
  
The name of Princeton University or Princeton may not be used in  
advertising or publicity pertaining to distribution of the software  
and/or database.  Title to copyright in this software, database and  
any associated documentation shall at all times remain with  
Princeton University and LICENSEE agrees to preserve same.  
```

### Used files

| Project file                   | Source file            |
|:------------------------------:|:----------------------:|
| `res/adj` &rarr; `embed/adj`   | `data.adj`             |
| `res/adv` &rarr; `embed/adv`   | `data.adv`             |
| `res/noun` &rarr; `embed/noun` | `data.noun`            |
| `res/verb` &rarr; `embed/verb` | `data.verb`            |
| `res/adj.irr`, `res/adj.suf`   | `adj.exc`, `adv.exc`   |

### Modifications

`adj.irr` and `adj.suf` were manually constructed using their respective source files. Automated formatting via [scripts](internal/scripts/) was applied to the remaining source files. Below is the summary of the procedure:

1. [scripts/res](internal/scripts/res/res.go)
    1. Remove license text from the beginning of the file.
    2. Extract words from the surrounding WordNet metadata.
    3. Remove single-letter words.
    4. Remove words containing apostrophes.
    5. Remove compound words.
    6. Remove words containing numbers.
    7. Remove entries consisting of multiple words.
    8. Remove proper nouns and adjectives derived from them.
    9. Remove parenthesized line content.
    10. If any [irregular verb](res/verb.irr) is missing from the main list, append it.
    11. Remove duplicate words.
    12. Remove unsuitable words. The lists of excluded words are available in [res/filters](res/filters/) directory.
    13. Change spelling of the selected words. To review the modifications, refer to [res/misc/replacements.json](res/misc/replacements.json).
    14. Sort word lists alphabetically.
2. [scripts/embed](internal/scripts/embed/embed.go)
    1. Append additional data to every word: its [FormType](formType.go#L4) and irregular forms (if applicable). Save newly compiled lists in the [embed](embed/) directory.
