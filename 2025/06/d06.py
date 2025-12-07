import re
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    lines, ops = content[:-1], re.split(r' +', content[-1])
    matrix = []
    for line in lines:
        digits = [int(d) for d in re.split(r' +', line.replace('\n', '').strip())]
        matrix.append(digits)
    result = 0
    for x in range(0, len(ops)):
        op = ops[x]
        if op == '+':
            sum = 0
            for y in range(0, len(matrix)):
                sum += matrix[y][x]
            result += sum
        elif op == '*':
            product = 1
            for y in range(0, len(matrix)):
                product *= matrix[y][x]
            result += product
        else:
            print(op)
            raise NotImplemented
    return result


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    lines = [l.replace('\n', '') for l in content[:-1]]
    width = max(len(l) for l in lines)
    lines = [l.ljust(width, ' ') for l in lines]
    ops = re.split(r' +', content[-1])
    result = 0
    for op in reversed(ops):
        if op == '+':
            acc = 0
        elif op == '*':
            acc = 1
        else:
            print(op)
            raise NotImplemented
        ns = []
        while len(lines[0]) > 0:
            n = ''
            nxt = []
            for line in lines:
                nxt.append(line[:-1])
                n += line[-1]
            n = n.strip()
            if n == '':
                lines = nxt
                break
            ns.append(n)
            if op == '+':
                acc += int(n)
            elif op == '*':
                acc *= int(n)
            lines = nxt
        result += acc
    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
