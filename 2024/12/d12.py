import sys
from pathlib import Path

sys.path.append(str(Path(__file__).parent.parent.parent))
from util.d2 import neighbours


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    height, width = len(content), len(content[0].rstrip())
    visited = {y: {x: False for x in range(width)} for y in range(height)}
    values = {y: {x: c for x, c in enumerate(line.rstrip())} for y, line in enumerate(content)}
    result = 0
    for y in range(height):
        for x in range(width):
            if not visited[y][x]:
                visited[y][x] = True
                plot = expand(width, height, values, visited, {(x, y)})
                result += price(plot, width, height)

    return result


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    height, width = len(content), len(content[0].rstrip())
    visited = {y: {x: False for x in range(width)} for y in range(height)}
    values = {y: {x: c for x, c in enumerate(line.rstrip())} for y, line in enumerate(content)}
    result = 0
    for y in range(height):
        for x in range(width):
            if not visited[y][x]:
                visited[y][x] = True
                plot = expand(width, height, values, visited, {(x, y)})
                result += price2(plot, values)

    return result


def price(plot, width, height):
    area = len(plot)
    perimeter = 0
    for x, y in plot:
        for dx, dy in neighbours():
            nx, ny = x + dx, y + dy
            if ny < 0 or ny >= height or nx < 0 or nx >= width or (nx, ny) not in plot:
                perimeter += 1
    return area * perimeter


def price2(plot, values):
    area = len(plot)
    vertices = 0
    for x, y in plot:
        value = values[y][x]
        left = values[y][x - 1] if x - 1 in values[y] else '.'
        right = values[y][x + 1] if x + 1 in values[y] else '.'
        up = values[y - 1][x] if y - 1 in values else '.'
        down = values[y + 1][x] if y + 1 in values else '.'

        left_up = values[y - 1][x - 1] if y - 1 in values and x - 1 in values[y - 1] else '.'
        if left_up != value and up == value and left == value:
            # inverted left-up
            vertices += 1
        if up != value and left != value:
            # left-up
            vertices += 1

        left_down = values[y + 1][x - 1] if y + 1 in values and x - 1 in values[y + 1] else '.'
        if left_down != value and down == value and left == value:
            # inverted left-down
            vertices += 1
        if down != value and left != value:
            # left-down
            vertices += 1

        right_up = values[y - 1][x + 1] if y - 1 in values and x + 1 in values[y - 1] else '.'
        if right_up != value and up == value and right == value:
            # inverted right-up
            vertices += 1
        if up != value and right != value:
            # right-up
            vertices += 1

        right_down = values[y + 1][x + 1] if y + 1 in values and x + 1 in values[y + 1] else '.'
        if right_down != value and down == value and right == value:
            # inverted right-down
            vertices += 1
        if down != value and right != value:
            # right-down
            vertices += 1

    return area * vertices


def expand(width, height, squares, visited, plot):
    current = set(plot)
    while len(current) > 0:
        nxt = set()
        for x, y in current:
            for dx, dy in neighbours():
                nx, ny = x + dx, y + dy
                if 0 <= ny < height and 0 <= nx < width and \
                        not visited[ny][nx] and \
                        squares[ny][nx] == squares[y][x]:
                    visited[ny][nx] = True
                    if (nx, ny) not in plot:
                        plot.add((nx, ny))
                        nxt.add((nx, ny))
        current = set(nxt)
    return plot


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
