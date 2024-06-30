#!/usr/bin/env python3

"""
This script extracts words from a WordNet database file
and formats them for use with neng.
"""

from argparse         import ArgumentParser
from better_profanity import profanity
from os.path          import exists
from re               import findall
from sys              import exit


RES_DIR        = "res"
FILTER_DIR     = RES_DIR + "/filters"
ORIG_DIR       = RES_DIR + "/wordnet"

LICENSE_OFFSET = 29
WORDS_COLUMN   =  4

SOURCE_FILES   = ("data.adj", "data.adv", "data.noun", "data.verb")
VERB_IRR_FILE  = "verb.irr"

REPLACEMENTS   = {
    "noun": {
        "adz"   : "adze",
        "ax"    : "axe",
        "cutlas": "cutlass",
        "poleax": "poleaxe"
    },
    "verb": {
        "poleax": "poleaxe"
    }
}


parser = ArgumentParser(
    prog="build_res.py",
    description="This script formats WordNet files for neng. Run it in neng's root directory."
)

parser.add_argument("-f", "--force", action="store_true", help="Overwrite a resource file if it exists.")
parser.add_argument("-g", "--generate-filters", action="store_true", help="Generate mature language filters instead of resource files.")
parser.add_argument("-m", "--filter-mature", action="store_true", help="Generate resource files, filtering it against filter files.")

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


def censor_lines(lines: [str], censored: [str]) -> [str]:
    """Returns lines filtered against censored list."""

    filtered = []
    for ln in lines:
        if ln not in censored:
            filtered.append(ln)
    
    return filtered


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


def get_mature_language(lines: [str]) -> [str]:
    """Returns mature language found in lines."""

    censored = []
    for ln in lines:
        if profanity.contains_profanity(ln):
            censored.append(ln)

    return censored


def load_file(path: str) -> [str]:
    """Loads lines from the file at path."""

    try:
        with open(path, mode='r') as f:
            return [ln.strip('\r').strip('\n') for ln in f.readlines()]
    except FileNotFoundError:
        print(f"{path} does not exist")
        exit(1)


def load_filter_file(path: str) -> [str]:
    """
    Attempts to load filter at path. If it does not exist, attempts to load automatically generated filter file ('path.auto').
    If filter is found, its content is returned. Otherwise, the scripts terminates with a message.
    """

    lines = []
    
    try:
        f = open(path, mode='r')
        print(f"found {path}. Applying.")
    except FileNotFoundError:
        try:
            f = open(f"{path}.auto", mode='r')
            print(f"found {path}.auto. Applying.")
        except FileNotFoundError:
            print(f"Filter '{path}' not found. Was it generated?")
            exit(1)

    lines = [ln.strip('\r').strip('\n') for ln in f.readlines()]
    f.close()

    return lines


def modify_list(fname: str, lines: [str]) -> [str]:
    """Performs file-specific modifications, mainly spelling changes."""

    if fname in REPLACEMENTS:
        for old, new in REPLACEMENTS[fname].items():
            try:
                lines[lines.index(old)] = new
            except ValueError as e:
                print(f"modify_list: {e} '{fname}'")
                continue

    return lines


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
        path      = f"{ORIG_DIR}/{file}"
        new_fname = file.split('.')[1]
        new_path  = f"{RES_DIR}/{new_fname}"
        
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
            lines = append_missing_verbirr(lines)

        lines = modify_list(new_fname, lines)

        if args.generate_filters:
            censored = get_mature_language(lines)

            fname = f"{FILTER_DIR}/{new_fname}.filter.auto"
            write_file(fname, censored)

            print(f"'{fname}' generated. Review and rename it '{fname.strip(".auto")}' or leave it as is and run the script again with '-m' to apply it.")
        elif args.filter_mature:
            censored = load_filter_file(f"{FILTER_DIR}/{new_fname}.filter")
            lines    = censor_lines(lines, censored)

            write_file(new_path, lines)
        else:
            write_file(new_path, lines)
