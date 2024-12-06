from collections import defaultdict
from pathlib import Path


def middle(xs):
    m = float(len(xs)) / 2
    if m % 2 != 0:
        return int(xs[int(m - .5)])
    else:
        raise NotImplementedError(m)


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    rules = defaultdict(lambda: [])
    result = 0
    for line in content:
        if '|' in line:
            a, b = line.rstrip().split('|')
            rules[b].append(a)
        if ',' in line:
            elements = line.rstrip().split(',')
            if is_valid(elements, rules):
                result += middle(elements)
    return result


def is_valid(elements: list[str], rules: dict[str, list[str]]) -> bool:
    stack = set()
    for element in elements:
        if element in rules:
            for expected in rules[element]:
                stack.add(expected)
        if element in stack:
            return False
    return True


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    rules = defaultdict(lambda: [])
    result = 0
    for line in content:
        if '|' in line:
            a, b = line.rstrip().split('|')
            rules[b].append(a)
        if ',' in line:
            elements = line.rstrip().split(',')
            if not is_valid(elements, rules):
                while not is_valid(elements, rules):
                    stack = set()
                    for i, element in enumerate(elements):
                        if element in rules:
                            for expected in rules[element]:
                                stack.add(expected)
                        if element in stack:
                            elements[i], elements[i-1] = elements[i-1], elements[i]
                            break
                result += middle(elements)

    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
