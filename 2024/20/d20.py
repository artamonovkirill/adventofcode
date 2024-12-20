from collections import defaultdict
from pathlib import Path

from util.d2 import neighbours


def solve(file: str) -> dict:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()

    walls = set()
    start = None, None
    end = None
    width = len(content[0].rstrip())
    height = len(content)
    for y, line in enumerate(content):
        for x, c in enumerate(line.rstrip()):
            if c == '#':
                walls.add((x, y))
            elif c == 'S':
                start = x, y
            elif c == 'E':
                end = x, y
    assert start
    assert end

    costs = dict()
    for ref in [start, end]:
        costs[ref] = {ref: 0}
        current = {ref}
        while len(current) > 0:
            nxt = set()
            for x, y in current:
                for dx, dy in neighbours():
                    nx = x + dx
                    ny = y + dy
                    if (nx, ny) not in walls:
                        cost = costs[ref][(x, y)] + 1
                        if (nx, ny) not in costs[ref] or costs[ref][(nx, ny)] > cost:
                            costs[ref][(nx, ny)] = cost
                            nxt.add((nx, ny))
            current = nxt

    assert costs[start][end] == costs[end][start]
    reference = costs[start][end]

    for y in range(height):
        for x in range(width):
            if (x, y) not in walls and (x + 1, y) in walls and (x + 2, y) in walls and (x + 3, y) not in walls:
                raise NotImplementedError()
            if (x, y) not in walls and (x, y + 1) in walls and (x, y + 2) in walls and (x, y + 3) not in walls:
                raise NotImplementedError()

    cheats = defaultdict(lambda: 0)
    for y in range(height):
        for x in range(1, width - 1):
            if (x - 1, y) not in walls and (x, y) in walls and (x + 1, y) not in walls:
                lr = costs[start][(x - 1, y)] + 2 + costs[end][(x + 1, y)]
                if lr < reference:
                    cheats[reference - lr] += 1
                rl = costs[start][(x + 1, y)] + 2 + costs[end][(x - 1, y)]
                if rl < reference:
                    cheats[reference - rl] += 1

    for y in range(1, height - 1):
        for x in range(width):
            if (x, y - 1) not in walls and (x, y) in walls and (x, y + 1) not in walls:
                td = costs[start][(x, y - 1)] + 2 + costs[end][(x, y + 1)]
                if td < reference:
                    cheats[reference - td] += 1
                dt = costs[start][(x, y + 1)] + 2 + costs[end][(x, y - 1)]
                if dt < reference:
                    cheats[reference - dt] += 1

    return dict(cheats)


if __name__ == "__main__":
    cheats = solve('puzzle.txt')
    print(sum(cheats[c] for c in cheats if c >= 100))
