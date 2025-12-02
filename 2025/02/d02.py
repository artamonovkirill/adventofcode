import re
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.read().split(',')

    result = 0
    for item in content:
        start, end = item.split('-')
        if start.startswith('0') or end.startswith('0'):
            raise NotImplemented
        for i in range(int(start), int(end) + 1):
            if re.match(r'^([0-9]+)\1$', str(i)):
                result += i
    return result


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.read().split(',')

    result = 0
    for item in content:
        start, end = item.split('-')
        if start.startswith('0') or end.startswith('0'):
            raise NotImplemented
        for i in range(int(start), int(end) + 1):
            if re.match(r'^([0-9]+)\1+$', str(i)):
                result += i
    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
