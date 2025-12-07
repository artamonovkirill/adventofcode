from collections import defaultdict
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = [l.replace('\n', '') for l in f.readlines()]
    for x in range(0, len(content[0])):
        if content[0][x] == 'S':
            xs = {x}
            break
    splits = 0
    for y in range(1, len(content)):
        next_xs = set()
        for x in xs:
            if content[y][x] == '^':
                splits += 1
                next_xs.add(x - 1)
                next_xs.add(x + 1)
            else:
                next_xs.add(x)
        xs = next_xs
    return splits


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = [l.replace('\n', '') for l in f.readlines()]
    for x in range(0, len(content[0])):
        if content[0][x] == 'S':
            xs = {x: 1}
            break
    for y in range(1, len(content)):
        next_xs = defaultdict(lambda: 0)
        for x, count in xs.items():
            if content[y][x] == '^':
                next_xs[x - 1] += count
                next_xs[x + 1] += count
            else:
                next_xs[x] += count
        xs = next_xs
    result = 0
    for count in next_xs.values():
        result += count
    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
