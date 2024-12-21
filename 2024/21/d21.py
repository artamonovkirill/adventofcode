import re
from pathlib import Path

numeric_keys = [
    ['7', '8', '9'],
    ['4', '5', '6'],
    ['1', '2', '3'],
    [None, '0', 'A']
]

directional_keys = [
    [None, '^', 'A'],
    ['<', 'v', '>']
]

directions = [
    (-1, 0, '<'),
    (1, 0, '>'),
    (0, -1, '^'),
    (0, 1, 'v')
]


def init(keys: list) -> dict:
    height = len(keys)
    width = len(keys[0])
    costs = dict()
    paths = dict()
    for y in range(height):
        for x in range(width):
            p = keys[y][x]
            if p is not None:
                costs[p] = {p: 0}
                paths[p] = {p: ['']}
                current = {(x, y)}
                while len(current) > 0:
                    nxt = set()
                    for cx, cy in current:
                        c = keys[cy][cx]
                        for dx, dy, symbol in directions:
                            nx, ny = cx + dx, cy + dy
                            if 0 <= nx < width and 0 <= ny < height and keys[ny][nx] is not None:
                                n = keys[ny][nx]
                                cost = costs[p][c] + 1
                                if n not in costs[p] or costs[p][n] > cost:
                                    costs[p][n] = cost
                                    paths[p][n] = [pp + symbol for pp in paths[p][c]]
                                    nxt.add((nx, ny))
                                elif costs[p][n] == cost:
                                    paths[p][n] += [pp + symbol for pp in paths[p][c]]
                    current = nxt
    return paths


numeric_keyboard = init(numeric_keys)
directional_keyboard = init(directional_keys)


def type_on(keyboard: dict, code: str) -> set:
    current = code[0]
    sequence = code[0]
    paths = {p + 'A' for p in keyboard['A'][current]}
    for c in code[1:]:
        next_paths = set()
        for path in paths:
            for connection in keyboard[current][c]:
                next_path = path + connection + 'A'
                next_paths.add(next_path)
        paths = next_paths
        sequence += c
        current = c
    return paths


directional_cache = dict()


def type_directions(code: str):
    if code not in directional_cache:
        paths = type_on(directional_keyboard, code)
        value = []
        for p in paths:
            matches = re.findall(r'[^A]*A', p)
            assert len(matches) > 0
            value.append(matches)
            directional_cache[code] = value
    return directional_cache[code]


sequences = [
    'A',
    '<A', '>A', 'vA', '^A',
    # three
    'v<A', 'v>A', 'vvA',
    '<vA', '<^A', '<^A', '<<A',
    '^>A', '^<A', '^^A',
    '>^A', '>vA', '>>A',
    # fours
    '^^^A', '^>^>A', '^>>^A', '^<<^A', '^<<A', '^^>A', '^>^A',
    '<v<A', '<^<A',
    '>>^A', '>^>A', '>vvA', '>>vA', '>v>A', '>^^A',
    'v<<A', 'vv>A' 'v>vA', 'vvvA', 'v>vA', 'vv>A', 'v>>A',
    # fives
    '^^>>A', '>>^^A', '>^>^A', '<^<^A', '<^^<A', '^^<<A', '^<^<A', '>^^>A', '<<^^A',
    # sixes
    '<^<^^A', '<^^<^A', '<^^^<A',
    '^^^<<A', '^<^<^A', '^^<^<A', '^<<^^A', '^<^^<A', '^^<<^A']


def solve(file: str, n: int = 25) -> int:
    cache = {0: dict()}
    for sequence in sequences:
        cache[0][sequence] = len(sequence)
    for i in range(1, n + 1):
        cache[i] = dict()
        for sequence in sequences:
            alternatives = type_directions(sequence)
            for alternative in alternatives:
                size = sum(cache[i - 1][s] for s in alternative)
                if sequence not in cache[i] or size < cache[i][sequence]:
                    cache[i][sequence] = size

    complexity = 0

    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    for line in content:
        code = line.rstrip()
        best = None
        paths = type_on(numeric_keyboard, code)
        for p in paths:
            size = sum(cache[n][m] for m in re.findall(r'[^A]+A', p))
            if best is None or size < best:
                best = size
        complexity += best * int(line[0:3])

    return complexity


if __name__ == "__main__":
    print(solve('puzzle.txt', 2))
    print(solve('puzzle.txt', 25))
