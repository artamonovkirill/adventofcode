from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    lines = [l.replace('\n', '') for l in content]

    papers = dict()
    width = len(lines[0])
    height = len(lines)
    for y in range(len(lines)):
        for x in range(len(lines[y])):
            if lines[y][x] == '.':
                papers[(x, y)] = False
            elif lines[y][x] == '@':
                papers[(x, y)] = True
            else:
                raise NotImplemented
    result = dict()
    for y in range(len(lines)):
        for x in range(len(lines[y])):
            if papers[(x, y)]:
                occupied = 0
                for dx, dy in neighbours():
                    nx = x + dx
                    ny = y + dy
                    if 0 <= nx < width and 0 <= ny < height and papers[(nx, ny)]:
                        occupied += 1
                if occupied < 4:
                    result[(x, y)] = True
    return len(result)


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    lines = [l.replace('\n', '') for l in content]
    papers = dict()
    width = len(lines[0])
    height = len(lines)
    for y in range(len(lines)):
        for x in range(len(lines[y])):
            if lines[y][x] == '.':
                papers[(x, y)] = False
            elif lines[y][x] == '@':
                papers[(x, y)] = True
            else:
                raise NotImplemented

    current = find_once(width, height, papers)
    result = 0
    while len(current) > 0:
        result += len(current)
        for x, y in current:
            papers[(x, y)] = False
        current = find_once(width, height, papers)
    return result


def find_once(width: int, height: int, papers: dict) -> dict:
    result = dict()
    for y in range(height):
        for x in range(width):
            if papers[(x, y)]:
                occupied = 0
                for dx, dy in neighbours():
                    nx = x + dx
                    ny = y + dy
                    if 0 <= nx < width and 0 <= ny < height and papers[(nx, ny)]:
                        occupied += 1
                if occupied < 4:
                    result[(x, y)] = True
    return result


def neighbours() -> list[tuple[int, int]]:
    return [
        (-1, -1),
        (0, -1),
        (1, -1),
        (-1, 0),
        # (0, 0),
        (1, 0),
        (-1, 1),
        (0, 1),
        (1, 1),
    ]


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
