import re
from pathlib import Path


def count_coins(machine, x, y) -> 0:
    # ax * ai + bx * bi = x
    # ay * ai + by * bi = y
    (ax, ay), (bx, by) = machine
    assert ax != ay
    #  ay * ax * ai + ay * bx * bi =  ay * x
    # -ax * ay * ai - ax * by * bi = -ax * y
    # bi * (ay * bx - ax * by) = ay * x - ax * y
    if (ay * x - ax * y) % (ay * bx - ax * by) == 0:
        bi = (ay * x - ax * y) // (ay * bx - ax * by)
        if (x - bx * bi) % ax == 0:
            ai = (x - bx * bi) // ax
            return ai * 3 + bi

    return 0


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()

    machine = None
    result = 0
    for line in content:
        if line.startswith('Button A:'):
            match = re.search(r'.*X\+([0-9]+), Y\+([0-9]+).*', line)
            machine = [(int(match.group(1)), int(match.group(2)))]
        elif line.startswith('Button B:'):
            match = re.search(r'.*X\+([0-9]+), Y\+([0-9]+).*', line)
            machine.append((int(match.group(1)), int(match.group(2))))
        elif line.startswith('Prize:'):
            match = re.search(r'.*X=([0-9]+), Y=([0-9]+).*', line)
            result += count_coins(machine, int(match.group(1)), int(match.group(2)))

    return result


def solve2(file: str) -> list[bool]:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()

    machine = None
    result = []
    for line in content:
        if line.startswith('Button A:'):
            match = re.search(r'.*X\+([0-9]+), Y\+([0-9]+).*', line)
            machine = [(int(match.group(1)), int(match.group(2)))]
        elif line.startswith('Button B:'):
            match = re.search(r'.*X\+([0-9]+), Y\+([0-9]+).*', line)
            machine.append((int(match.group(1)), int(match.group(2))))
        elif line.startswith('Prize:'):
            match = re.search(r'.*X=([0-9]+), Y=([0-9]+).*', line)
            result.append(
                count_coins(machine, 10000000000000 + int(match.group(1)), 10000000000000 + int(match.group(2))))

    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(sum(solve2('puzzle.txt')))
