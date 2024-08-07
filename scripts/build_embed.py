#!/usr/bin/env python3

"""
This script builds the embedded resource for neng 
from the lists built by build_res.py.
"""

import enum
import utils


class W_TYPE(enum.Enum):
    """Corresponds to neng.WordType type."""

    REGULAR      = 0
    IRREGULAR    = 1
    PLURAL_ONLY  = 2
    SUFFIXED     = 3
    UNCOMPARABLE = 4
    UNCOUNTABLE  = 5


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
        t: W_TYPE

        if n in irr:
            t = W_TYPE.IRREGULAR
        elif n in plo:
            t = W_TYPE.PLURAL_ONLY
        elif n in unc:
            t = W_TYPE.UNCOUNTABLE
        else:
            t = W_TYPE.REGULAR

        embed.append(f"{n} {int(t.value)}{' ' + irr[n] if t == W_TYPE.IRREGULAR else ''}")

    utils.write_file(f"{EMBED_DIR}/noun", embed)

def build_verb():
    """Builds the embedded "verb" file."""

    verb = utils.load_file(f"{RES_DIR}/verb")
    irr  = _load_irregular("verb.irr")

    embed = []
    for v in verb:
        t = W_TYPE.IRREGULAR if v in irr else W_TYPE.REGULAR
        embed.append(f"{v} {int(t.value)}{' ' + irr[v] if t == W_TYPE.IRREGULAR else ''}")

    utils.write_file(f"{EMBED_DIR}/verb", embed)


def _build_modifier(words_fname: str):
    """A helper function for building "adj" and "adv" files, which have the same structure."""

    words = utils.load_file(f"{RES_DIR}/{words_fname}")
    irr   = _load_irregular(f"{words_fname}.irr")
    ncmp  = utils.load_file(f"{RES_DIR}/{words_fname}.ncmp")
    suf   = utils.load_file(f"{RES_DIR}/{words_fname}.suf")

    embed = []
    for w in words:
        t: W_TYPE

        if w in irr:
            t = W_TYPE.IRREGULAR
        elif w in ncmp:
            t = W_TYPE.UNCOMPARABLE
        elif w in suf:
            t = W_TYPE.SUFFIXED
        else:
            t = W_TYPE.REGULAR

        embed.append(f"{w} {int(t.value)}{' ' + irr[w] if t == W_TYPE.IRREGULAR else ''}")

    utils.write_file(f"{EMBED_DIR}/{words_fname}", embed)


def _load_irregular(fname: str) -> {}:
    """Returns a dictionary containing irregular words. Base form words are the keys."""

    irr = {}
    for w in utils.load_file(f"{RES_DIR}/{fname}"):
        s = w.split(',', maxsplit=1)
        irr[s[0]] = s[1]

    return irr


if __name__ == "__main__":
    build_adj()
    build_adv()
    build_noun()
    build_verb()