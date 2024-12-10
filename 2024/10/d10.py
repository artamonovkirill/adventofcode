import sys
from collections import defaultdict
from pathlib import Path

sys.path.append(str(Path(__file__).parent.parent.parent))
from util.d2 import neighbours


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    points = defaultdict(lambda: [])
    for y, line in enumerate(content):
        for x, c in enumerate(line.rstrip()):
            if c != '.':
                points[int(c)].append((x, y))
    current = {p: {p} for p in points[9]}
    for i in reversed(range(9)):
        nxt = dict()
        for x, y in points[i]:
            peaks = set()
            for dx, dy in neighbours():
                if (x + dx, y + dy) in current:
                    for p in current[(x + dx, y + dy)]:
                        peaks.add(p)
            nxt[(x, y)] = peaks
        current = nxt
    return sum(len(current[k]) for k in current)


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    points = defaultdict(lambda: [])
    for y, line in enumerate(content):
        for x, c in enumerate(line.rstrip()):
            if c != '.':
                points[int(c)].append((x, y))
    current = {p: 1 for p in points[9]}
    for i in reversed(range(9)):
        nxt = dict()
        for x, y in points[i]:
            count = 0
            for dx, dy in neighbours():
                if (x + dx, y + dy) in current:
                    count += current[(x + dx, y + dy)]
            nxt[(x, y)] = count
        current = nxt
    return sum(current[k] for k in current)


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
