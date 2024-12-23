import itertools
from collections import defaultdict
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    connections = defaultdict(lambda: set())
    ts = set()
    for line in content:
        connection = line.rstrip()
        a, b = connection.split('-')
        connections[a].add(b)
        connections[b].add(a)
        if a.startswith('t'):
            ts.add(a)
        if b.startswith('t'):
            ts.add(b)

    trios = set()
    for t in ts:
        if len(connections[t]) > 1:
            for c in itertools.combinations(connections[t], 2):
                if c[1] in connections[c[0]]:
                    l = sorted([t, c[0], c[1]])
                    trios.add((l[0], l[1], l[2]))

    return len(trios)


if __name__ == "__main__":
    print(solve('puzzle.txt'))
