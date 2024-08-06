def filter_containing(lines: [str], char: str) -> [str]:
    """Removes the lines containing char."""

    filtered = []
    for ln in lines:
        if char not in ln:
            filtered.append(ln)

    return filtered
