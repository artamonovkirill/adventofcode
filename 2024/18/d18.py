import re
from pathlib import Path

from util.d2 import neighbours


def solve(file: str, n: int, size: int) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    bytes = set()
    for line in content:
        match = re.search(r'([0-9]+),([0-9]+)', line).groups()
        assert len(match) == 2
        bytes.add((int(match[0]), int(match[1])))
        if len(bytes) >= n:
            break

    current = {(0, 0)}
    steps = {(0, 0): 0}
    while len(current) > 0:
        nxt = set()
        for x, y in current:
            for dx, dy in neighbours():
                next_x, next_y = x + dx, y + dy
                if 0 <= next_y < size and 0 <= next_x < size and (next_x, next_y) not in bytes:
                    next_steps = steps[(x, y)] + 1
                    if (next_x, next_y) not in steps or next_steps < steps[(next_x, next_y)]:
                        nxt.add((next_x, next_y))
                        steps[(next_x, next_y)] = next_steps
        current = nxt

    return steps[(size - 1, size - 1)]


def solve2(file: str, n: int, size: int) -> tuple[int, int]:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    all_bytes = []
    for line in content:
        match = re.search(r'([0-9]+),([0-9]+)', line).groups()
        assert len(match) == 2
        all_bytes.append((int(match[0]), int(match[1])))

    for i in range(n + 1, len(all_bytes) + 1):
        blocked = set(all_bytes[:i])
        current = {(0, 0)}
        steps = {(0, 0): 0}
        while len(current) > 0:
            nxt = set()
            for x, y in current:
                for dx, dy in neighbours():
                    next_x, next_y = x + dx, y + dy
                    if 0 <= next_y < size and 0 <= next_x < size and (next_x, next_y) not in blocked:
                        next_steps = steps[(x, y)] + 1
                        if (next_x, next_y) not in steps or next_steps < steps[(next_x, next_y)]:
                            nxt.add((next_x, next_y))
                            steps[(next_x, next_y)] = next_steps
            current = nxt
        if (size - 1, size - 1) in steps:
            continue
        else:
            return all_bytes[i - 1]

    raise NotImplementedError()


if __name__ == "__main__":
    print(solve('puzzle.txt', 1024, 71))
    print(solve2('puzzle.txt', 1024, 71))
