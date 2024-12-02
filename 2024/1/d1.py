from collections import defaultdict
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    lefts = []
    rights = []
    with path.open() as f:
        content = f.readlines()
    for line in content:
        left, right = line.split('   ')
        lefts.append(int(left))
        rights.append(int(right.strip()))
    lefts.sort()
    rights.sort()
    return sum(abs(l - r) for l, r in zip(lefts, rights))


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    lefts = []
    rights = defaultdict(lambda: 0)
    with path.open() as f:
        content = f.readlines()
    for line in content:
        left, right = line.split('   ')
        lefts.append(int(left))
        rights[int(right.strip())] += 1
    return sum(l * rights[l] for l in lefts)


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
