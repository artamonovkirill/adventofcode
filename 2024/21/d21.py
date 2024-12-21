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


def type_on(keyboard: dict, code: str) -> set[str]:
    current = code[0]
    sequence = code[0]
    paths = {p + 'A' for p in keyboard['A'][current]}
    for c in code[1:]:
        best = None
        for path in paths:
            for connection in keyboard[current][c]:
                next_path = path + connection + 'A'
                size = len(next_path)
                if best is None or best > size:
                    best = size
                    next_paths = {next_path}
                elif best == size:
                    next_paths.add(next_path)
        paths = next_paths
        sequence += c
        current = c
    return paths, best


def type_sequentially(code: str, keyboards: list[dict]) -> set[str]:
    paths, best = type_on(keyboards[0], code)
    for keyboard in keyboards[1:]:
        print(code, )
        best = None
        for path in paths:
            next_paths, size = type_on(keyboard, path)
            if best is None or best > size:
                best = size
                next_next_paths = next_paths
            elif best == size:
                for p in next_paths:
                    next_next_paths.add(p)
        paths = next_next_paths
    return paths, best


def solve(file: str) -> int:
    layers = [numeric_keyboard, directional_keyboard, directional_keyboard]

    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    complexity = 0
    for line in content:
        _, size = type_sequentially(line.rstrip(), layers)
        complexity += size * int(line[0:3])
    return complexity


if __name__ == "__main__":
    print(solve('puzzle.txt'))
