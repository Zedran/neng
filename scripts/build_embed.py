#!/usr/bin/env python3

"""
This script builds the embedded resource for neng 
from the lists built by build_res.py.
"""

import enum
import os
import sys
import utils


class FORM_TYPE(enum.Enum):
    """Corresponds to neng.FormType."""

    REGULAR        = 0
    IRREGULAR      = 1
    PLURAL_ONLY    = 2
    SUFFIXED       = 3
    NON_COMPARABLE = 4
    UNCOUNTABLE    = 5


EMBED_DIR      = "embed"
RES_DIR        = "res"


def build_adj():
    """Calls the generalized function that builds the embedded "adj" file."""

    _build_modifier("adj")


def build_adv():
    """Calls the generalized function that builds the embedded "adv" file."""

    _build_modifier("adv")


def build_noun():
    """Builds the embedded "noun" file."""

    noun = utils.load_file(f"{RES_DIR}/noun")
    irr  = _load_irregular("noun.irr")
    plo  = utils.load_file(f"{RES_DIR}/noun.plo")
    unc  = utils.load_file(f"{RES_DIR}/noun.unc")

    embed = []
    for n in noun:
        t: FORM_TYPE

        if n in irr:
            t = FORM_TYPE.IRREGULAR
        elif n in plo:
            t = FORM_TYPE.PLURAL_ONLY
        elif n in unc:
            t = FORM_TYPE.UNCOUNTABLE
        else:
            t = FORM_TYPE.REGULAR

        embed.append(_build_line(t, n, irr))

    utils.write_file(f"{EMBED_DIR}/noun", False, embed)

def build_verb():
    """Builds the embedded "verb" file."""

    verb = utils.load_file(f"{RES_DIR}/verb")
    irr  = _load_irregular("verb.irr")

    embed = []
    for v in verb:
        t = FORM_TYPE.IRREGULAR if v in irr else FORM_TYPE.REGULAR
        embed.append(_build_line(t, v, irr))

    utils.write_file(f"{EMBED_DIR}/verb", False, embed)


def _build_line(t: FORM_TYPE, word: str, irr: [str]) -> str:
    """A helper function that builds a single line of the embedded file."""

    return f"{int(t.value)}{word}{',' + irr[word] if t == FORM_TYPE.IRREGULAR else ''}"


def _build_modifier(words_fname: str):
    """A helper function for building "adj" and "adv" files, which have the same structure."""

    words = utils.load_file(f"{RES_DIR}/{words_fname}")
    irr   = _load_irregular(f"{words_fname}.irr")
    ncmp  = utils.load_file(f"{RES_DIR}/{words_fname}.ncmp")
    suf   = utils.load_file(f"{RES_DIR}/{words_fname}.suf")

    embed = []
    for w in words:
        t: FORM_TYPE

        if w in irr:
            t = FORM_TYPE.IRREGULAR
        elif w in ncmp:
            t = FORM_TYPE.NON_COMPARABLE
        elif w in suf:
            t = FORM_TYPE.SUFFIXED
        else:
            t = FORM_TYPE.REGULAR

        embed.append(_build_line(t, w, irr))

    utils.write_file(f"{EMBED_DIR}/{words_fname}", False, embed)


def _load_irregular(fname: str) -> {}:
    """Returns a dictionary containing irregular words. Base form words are the keys."""

    irr = {}
    for w in utils.load_file(f"{RES_DIR}/{fname}"):
        s = w.split(',', maxsplit=1)
        irr[s[0]] = s[1]

    return irr


if __name__ == "__main__":
    try:
        os.mkdir(EMBED_DIR)
    except FileExistsError:
        pass
    except Exception as e:
        print(e)
        sys.exit(1)

    build_adj()
    build_adv()
    build_noun()
    build_verb()
