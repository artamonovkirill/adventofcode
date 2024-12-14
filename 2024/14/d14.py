import re
from pathlib import Path


def solve(file: str, width: int, height: int) -> int:
    robots = parse(file)
    return score(robots, width, height, 100)


def parse(file: str):
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    robots = []
    for line in content:
        match = re.search(r'p=([0-9]+),([0-9]+) v=(-?[0-9]+),(-?[0-9]+)', line).groups()
        assert len(match) == 4
        robots.append([int(m) for m in match])
    return robots


def score(robots, width, height, i) -> int:
    positions = dict()
    for x, y, dx, dy in robots:
        x = (x + dx * i) % width
        y = (y + dy * i) % height
        if (x, y) in positions:
            positions[(x, y)] += 1
        else:
            positions[(x, y)] = 1
    q1, q2, q3, q4 = 0, 0, 0, 0
    mid_x = width // 2
    mid_y = height // 2

    for x, y in positions:
        if x < mid_x:
            if y < mid_y:
                q1 += positions[(x, y)]
            elif y > mid_y:
                q3 += positions[(x, y)]
        elif x > mid_x:
            if y < mid_y:
                q2 += positions[(x, y)]
            elif y > mid_y:
                q4 += positions[(x, y)]
    return q1 * q2 * q3 * q4


def solve2(file: str, width: int, height: int) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    all_positions = []
    robots = []
    for line in content:
        match = re.search(r'p=([0-9]+),([0-9]+) v=(-?[0-9]+),(-?[0-9]+)', line).groups()
        assert len(match) == 4
        x, y, dx, dy = [int(m) for m in match]
        robots.append((x, y, dx, dy))
        positions = dict()
        i = 0
        while (x, y) not in positions:
            positions[(x, y)] = i
            x = (x + dx) % width
            y = (y + dy) % height
            i += 1
        all_positions.append(positions)

    assert len({len(p) for p in all_positions}) == 1
    cycle = len(all_positions[0])

    line_length = 10
    for i in range(cycle + 1):
        positions = set()
        for x, y, dx, dy in robots:
            x = (x + dx * i) % width
            y = (y + dy * i) % height
            positions.add((x, y))

        for y in range(height):
            for x in range(width - line_length + 1):
                if all((x + dx, y) in positions for dx in range(line_length)):
                    print(i)
                    print('\n'.join(
                        ''.join(
                            '#' if (x, y) in positions else ' '
                            for x in range(width))
                        for y in range(height)))
                    break


if __name__ == "__main__":
    print(solve('puzzle.txt', 101, 103))
    print(solve2('puzzle.txt', 101, 103))
