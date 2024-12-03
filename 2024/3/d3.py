import re
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    result = 0
    for line in content:
        matches = re.findall(r'mul[(]([0-9]{1,3}),([0-9]{1,3})[)]', line)
        for (a, b) in matches:
            result += int(a) * int(b)
    return result


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    result = 0
    do = True
    for line in content:
        matches = re.findall(r"(mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\))", line)
        for match in matches:
            if match[0] == 'do()':
                do = True
            elif match[0] == "don't()":
                do = False
            elif do:
                _, a, b = match
                result += int(a) * int(b)
    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
