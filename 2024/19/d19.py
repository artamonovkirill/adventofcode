from collections import defaultdict
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    possible = 0
    for line in content:
        line = line.rstrip()
        if ',' in line:
            patterns = set(line.split(', '))
        elif line and is_possible(line, patterns):
            possible += 1
    return possible


def is_possible(line: str, patterns: list[str]) -> bool:
    current = {''}
    while len(current) > 0:
        nexts = set()
        for c in current:
            for p in patterns:
                nxt = c + p
                if line == nxt:
                    return True
                elif line.startswith(nxt):
                    nexts.add(nxt)
        current = nexts
    return False


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    possible = 0
    for line in content:
        line = line.rstrip()
        if ',' in line:
            patterns = set(line.split(', '))
        elif line:
            current = {'': 1}
            while len(current) > 0:
                nexts = defaultdict(lambda: 0)
                for c in current:
                    previous = current[c]
                    for p in patterns:
                        nxt = c + p
                        if line == nxt:
                            possible += previous
                        elif line.startswith(nxt):
                            nexts[nxt] += previous
                current = nexts
    return possible


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
