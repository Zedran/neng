#!/usr/bin/env python3

"""
This script extracts words from a WordNet database file
and formats them for use with neng.
"""

from argparse import ArgumentParser
from os.path  import exists
from re       import findall
from sys      import exit


RES_DIR        = "res"
ORIG_DIR       = RES_DIR + "/wordnet"

LICENSE_OFFSET = 29
WORDS_COLUMN   =  4

SOURCE_FILES   = ("data.adj", "data.noun", "data.verb")
VERB_IRR_FILE  = "verb.irr"


parser = ArgumentParser(
    prog="build_res.py",
    description="This script formats WordNet files for neng. Run in neng's root directory."
)

parser.add_argument("-f", "--force", action="store_true", help="Overwrite a resource file if it exists.")

args = parser.parse_args()


def append_missing_verbirr(lines: [str]) -> [str]:
    """Appends irregular verbs that are missing from the source verb file to the processed list."""

    with open(f"{RES_DIR}/{VERB_IRR_FILE}", mode='r') as ivf:
        verbirr = [ln.split(',')[0] for ln in ivf.readlines()]

        for iv in verbirr:
            if not iv in lines:
                lines.append(iv)   

    return lines


def filter_apostrophes(lines: [str]) -> [str]:
    """Removes words containing apostrophes."""

    filtered = []
    for ln in lines:
        if ln.find("'") == -1:
            filtered.append(ln)

    return filtered


def filter_compound_words(lines: [str]) -> [str]:
    """Removes compound words as they are mostly an adjective-noun pair."""

    filtered = []
    for ln in lines:
        if ln.find('-') == -1:
            filtered.append(ln)

    return filtered


def filter_duplicates(lines: [str]) -> [str]:
    """Removes duplicates, does not preserve order of elements."""
    
    return list(set(lines))


def filter_metadata(lines: [str]) -> [str]:
    """Gets a list of words in the fifth column of a file, discarding all other information."""

    filtered = []
    for ln in lines:
        s = ln.split(' ')
        if len(s[WORDS_COLUMN]) > 0:
            filtered.append(s[WORDS_COLUMN])

    return filtered


def filter_multiword_entries(lines: [str]) -> [str]:
    """
    Removes multi-word entries. Not all of them are suited for use
    in the generator and it is difficult to assess them automatically.
    """

    filtered = []
    for ln in lines:
        if ln.find('_') == -1:
            filtered.append(ln)

    return filtered


def filter_parentheses(lines: [str]) -> [str]:
    """Removes parentheses and content inside of them."""

    filtered = []

    for ln in lines:
        if ln.find('(') > 1:
            par = findall(r"\(.*?\)", ln)[0]
            filtered.append(ln.replace(par, ''))
        else:
            filtered.append(ln)
    
    return filtered


def filter_proper_nouns(lines: [str]) -> [str]:
    """Removes all proper nouns and adjectives derived from them."""

    filtered = []
    for ln in lines:
        if not ln[0].isupper():
            filtered.append(ln)

    return filtered


def filter_single_letter_words(lines: [str]) -> [str]:
    """Removes single-letter words and empty lines."""

    filtered = []
    for ln in lines:
        if len(ln) > 1:
            filtered.append(ln)
    
    return filtered


def filter_numbers(lines: [str]) -> [str]:
    """Removes words that contain numbers."""

    filtered = []
    for ln in lines:
        if not any(c.isdigit() for c in ln):
            filtered.append(ln)

    return filtered


def load_file(path: str) -> [str]:
    """Loads lines from the file at path."""

    try:
        with open(path, mode='r') as f:
            return f.readlines()
    except FileNotFoundError:
        print(f"{path} does not exist")
        exit(1)


def strip_license(lines: [str]) -> str:
    """Removes license at the beginning of a file."""

    return lines[LICENSE_OFFSET:]


def write_file(path: str, lines: [str]):
    """Sorts lines alphabetically and writes them to the file at path."""

    lines.sort()

    with open(path, mode='w') as f:
        f.write('\n'.join(lines))


if __name__ == "__main__":
    for file in SOURCE_FILES:
        path     = f"{ORIG_DIR}/{file}"
        new_path = f"{RES_DIR}/{file.split('.')[1]}"
        
        if exists(new_path) and not args.force:
            print(f"{new_path:<10} exists, skipping.")
            continue

        lines = load_file(path)

        lines = strip_license(lines)
        lines = filter_metadata(lines)
        lines = filter_single_letter_words(lines)

        lines = filter_multiword_entries(lines)
        lines = filter_proper_nouns(lines)
        lines = filter_parentheses(lines)
        lines = filter_duplicates(lines)
        lines = filter_compound_words(lines)
        lines = filter_numbers(lines)
        lines = filter_apostrophes(lines)

        if file == "data.verb":
            append_missing_verbirr(lines)

        write_file(new_path, lines)
