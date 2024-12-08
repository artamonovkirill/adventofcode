from collections import defaultdict
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    antennas = defaultdict(lambda: dict())
    height = len(content)
    width = len(content[0].rstrip())
    for y, line in enumerate(content):
        for x, c in enumerate(line.rstrip()):
            if c != '.':
                antennas[c][(x, y)] = True
    anti_nodes = set()
    for c in antennas:
        c_antennas = antennas[c]
        for y in range(height):
            for x in range(width):
                for (xa, ya) in c_antennas:
                    dx = xa - x
                    dy = ya - y
                    if (dx != 0 or dy != 0) and (x + 2 * dx, y + 2 * dy) in c_antennas:
                        anti_nodes.add((x, y))
    return len(anti_nodes)


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    antennas = defaultdict(lambda: dict())
    height = len(content)
    width = len(content[0].rstrip())
    steps = max(width, height)
    for y, line in enumerate(content):
        for x, c in enumerate(line.rstrip()):
            if c != '.':
                antennas[c][(x, y)] = True
    anti_nodes = set()
    for c in antennas:
        if len(antennas[c]) > 0:
            for a in antennas[c]:
                for a2 in antennas[c]:
                    if a != a2:
                        (x, y) = a
                        (x2, y2) = a2
                        dx = x - x2
                        dy = y - y2
                        for i in range(1, steps):
                            nxt_x = x2 + dx * i
                            nxt_y = y2 + dy * i
                            if nxt_x < 0 or nxt_x >= width or nxt_y < 0 or nxt_y >= height:
                                break
                            anti_nodes.add((nxt_x, nxt_y))
    return len(anti_nodes)


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
