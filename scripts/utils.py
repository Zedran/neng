def filter_containing(lines: [str], char: str) -> [str]:
    """Removes the lines containing char."""

    filtered = []
    for ln in lines:
        if char not in ln:
            filtered.append(ln)

    return filtered


def load_file(path: str) -> [str]:
    """Loads lines from the file at path."""

    try:
        with open(path, mode='r') as f:
            return [ln.strip('\r').strip('\n') for ln in f.readlines()]
    except FileNotFoundError:
        print(f"{path} does not exist")
        exit(1)


def write_file(path: str, sort: bool, lines: [str]):
    """Sorts lines alphabetically and writes them to the file at path."""

    if sort:
        lines.sort()

    with open(path, mode='w') as f:
        f.write('\n'.join(lines))
