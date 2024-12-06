from collections import defaultdict
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    obstacles = defaultdict(lambda: defaultdict(lambda: False))
    visited = defaultdict(lambda: defaultdict(lambda: False))
    guard = None
    height = len(content)
    width = len(content[0])
    for y, line in enumerate(content):
        for x, c in enumerate(line):
            if c == '#':
                obstacles[y][x] = True
            elif c == '^':
                guard = (x, y, 'north')
    while 0 <= guard[0] < width and 0 <= guard[1] < height:
        x, y, direction = guard
        visited[y][x] = True
        if direction == 'north':
            nx, ny = x, y - 1
            if obstacles[ny][nx]:
                guard = (x, y, 'east')
            else:
                guard = (nx, ny, 'north')
        elif direction == 'east':
            nx, ny = x + 1, y
            if obstacles[ny][nx]:
                guard = (x, y, 'south')
            else:
                guard = (nx, ny, 'east')
        elif direction == 'south':
            nx, ny = x, y + 1
            if obstacles[ny][nx]:
                guard = (x, y, 'west')
            else:
                guard = (nx, ny, 'south')
        elif direction == 'west':
            nx, ny = x - 1, y
            if obstacles[ny][nx]:
                guard = (x, y, 'north')
            else:
                guard = (nx, ny, 'west')
        else:
            raise NotImplementedError(guard)
    return sum(len(visited[v]) for v in visited)


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    obstacles = defaultdict(lambda: defaultdict(lambda: False))
    candidates = defaultdict(lambda: defaultdict(lambda: False))
    start = None
    height = len(content)
    width = len(content[0])
    for y, line in enumerate(content):
        for x, c in enumerate(line):
            if c == '#':
                obstacles[y][x] = True
            elif c == '^':
                start = (x, y, 'north')
            else:
                candidates[y][x] = True
    result = 0
    for cy in candidates:
        for cx in candidates[cy]:
            obstacles[cy][cx] = True
            guard = start
            result += loop(obstacles, guard, width, height)
            obstacles[cy][cx] = False
    return result


def loop(obstacles, guard, width, height):
    visited = defaultdict(lambda: defaultdict(lambda: []))
    while 0 <= guard[0] < width and 0 <= guard[1] < height:
        x, y, direction = guard
        if direction in visited[y][x]:
            return 1
        visited[y][x].append(direction)
        if direction == 'north':
            nx, ny = x, y - 1
            if obstacles[ny][nx]:
                guard = (x, y, 'east')
            else:
                guard = (nx, ny, 'north')
        elif direction == 'east':
            nx, ny = x + 1, y
            if obstacles[ny][nx]:
                guard = (x, y, 'south')
            else:
                guard = (nx, ny, 'east')
        elif direction == 'south':
            nx, ny = x, y + 1
            if obstacles[ny][nx]:
                guard = (x, y, 'west')
            else:
                guard = (nx, ny, 'south')
        elif direction == 'west':
            nx, ny = x - 1, y
            if obstacles[ny][nx]:
                guard = (x, y, 'north')
            else:
                guard = (nx, ny, 'west')
        else:
            raise NotImplementedError(guard)
    return 0


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
