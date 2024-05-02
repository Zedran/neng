# Notices and information

This software incorporates materials licensed by third parties.

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

| Project file | Source file |
|:------------:|:-----------:|
| `res/adj`    | `data.adj`  |
| `res/noun`   | `data.noun` |
| `res/verb`   | `data.verb` |

### Modifications

Source files underwent automated formatting with `build_res.py` script. Below is the summary of the procedure:

1. Remove license text from the beginning of the file
2. Extract words from the surrounding metadata
3. Remove single-letter words
4. Remove entries consisting of multiple words
5. Remove proper nouns and adjectives derived from them
6. Remove all of the content within the parentheses
7. Remove duplicate words
8. Remove compound words
9. Remove words containing numbers
10. Remove words containing apostrophes
11. Append missing irregular verbs (found in `res/verb.irr`) to the verb list
12. Filter mature and controversial language. Filter files were created with [better_profanity](https://github.com/snguyenthanh/better_profanity) Python package and then edited manually to include more words and remove false positives
13. Sort word lists alphabetically

The following manual modifications were carried out:

```text
noun: cutlas -> cutlass
```
