from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    result = 0
    for line in content:
        parts = line.rstrip().split(': ')
        expected = int(parts[0])
        values = [int(v) for v in parts[1].split(' ')]
        current = [values[0]]
        for v in values[1:]:
            nxt = []
            for c in current:
                sum = c + v
                if sum <= expected:
                    nxt.append(sum)
                product = c * v
                if product <= expected:
                    nxt.append(product)
            current = nxt
        if any(c == expected for c in current):
            result += expected
    return result


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    result = 0
    for line in content:
        parts = line.rstrip().split(': ')
        expected = int(parts[0])
        values = [int(v) for v in parts[1].split(' ')]
        current = [values[0]]
        for v in values[1:]:
            nxt = []
            for c in current:
                sum = c + v
                if sum <= expected:
                    nxt.append(sum)
                product = c * v
                if product <= expected:
                    nxt.append(product)
                concat = int(str(c) + str(v))
                if concat <= expected:
                    nxt.append(concat)
            current = nxt
        if any(c == expected for c in current):
            result += expected
    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
