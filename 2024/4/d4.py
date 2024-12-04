from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    lines = [l.rstrip() for l in content]
    height = len(lines)
    width = len(lines[0])
    count = 0
    for y in range(height):
        for x in range(width):
            if lines[y][x] == 'X':
                if x < width - 3 and \
                        lines[y][x + 1] == 'M' and \
                        lines[y][x + 2] == 'A' and \
                        lines[y][x + 3] == 'S':
                    count += 1
                if x > 2 and \
                        lines[y][x - 1] == 'M' and \
                        lines[y][x - 2] == 'A' and \
                        lines[y][x - 3] == 'S':
                    count += 1
                if y < height - 3 and \
                        lines[y + 1][x] == 'M' and \
                        lines[y + 2][x] == 'A' and \
                        lines[y + 3][x] == 'S':
                    count += 1
                if y > 2 and \
                        lines[y - 1][x] == 'M' and \
                        lines[y - 2][x] == 'A' and \
                        lines[y - 3][x] == 'S':
                    count += 1
                if x < width - 3 and y < height - 3 and \
                        lines[y + 1][x + 1] == 'M' and \
                        lines[y + 2][x + 2] == 'A' and \
                        lines[y + 3][x + 3] == 'S':
                    count += 1
                if x > 2 and y > 2 and \
                        lines[y - 1][x - 1] == 'M' and \
                        lines[y - 2][x - 2] == 'A' and \
                        lines[y - 3][x - 3] == 'S':
                    count += 1
                if x < width - 3 and y > 2 and \
                        lines[y - 1][x + 1] == 'M' and \
                        lines[y - 2][x + 2] == 'A' and \
                        lines[y - 3][x + 3] == 'S':
                    count += 1
                if x > 2 and y < height - 3 and \
                        lines[y + 1][x - 1] == 'M' and \
                        lines[y + 2][x - 2] == 'A' and \
                        lines[y + 3][x - 3] == 'S':
                    count += 1
    return count


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    lines = [l.rstrip() for l in content]
    height = len(lines)
    width = len(lines[0])
    count = 0
    for y in range(height - 2):
        for x in range(width - 2):
            if lines[y][x] == 'M' and lines[y + 1][x + 1] == 'A' and lines[y + 2][x + 2] == 'S':
                if lines[y + 2][x] == 'M' and lines[y][x + 2] == 'S':
                    count += 1
                if lines[y + 2][x] == 'S' and lines[y][x + 2] == 'M':
                    count += 1
            if lines[y][x] == 'S' and lines[y + 1][x + 1] == 'A' and lines[y + 2][x + 2] == 'M':
                if lines[y + 2][x] == 'M' and lines[y][x + 2] == 'S':
                    count += 1
                if lines[y + 2][x] == 'S' and lines[y][x + 2] == 'M':
                    count += 1
    return count


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
