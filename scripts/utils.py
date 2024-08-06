def filter_containing(lines: [str], char: str) -> [str]:
    """Removes the lines containing char."""

    filtered = []
    for ln in lines:
        if ln.find(char) == -1:
            filtered.append(ln)

    return filtered
