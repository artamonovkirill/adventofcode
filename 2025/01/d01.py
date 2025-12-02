from pathlib import Path

dial_size = 100


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()

    dial = 50
    zeros = 0
    for line in content:
        direction = line[0]
        steps = int(line[1:])
        if direction == 'L':
            dial = (dial + dial_size - steps) % dial_size
        elif direction == 'R':
            dial = (dial + steps) % dial_size
        else:
            raise NotImplemented("Unknown direction: " + direction)
        if dial == 0:
            zeros += 1
    return zeros


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()

    dial = 50
    zeros = 0
    for line in content:
        direction = line[0]
        steps = int(line[1:])

        zeros += steps // dial_size
        diff = steps % dial_size

        if direction == 'L':
            if diff > dial:
                if dial != 0:
                    zeros += 1
                dial += dial_size - diff
            elif diff == dial:
                zeros += 1
                dial = 0
            else:
                dial -= diff
        elif direction == 'R':
            if diff+dial > dial_size:
                zeros += 1
                dial += diff - dial_size
            elif diff+dial == dial_size:
                zeros += 1
                dial = 0
            else:
                dial += diff
        else:
            raise NotImplemented("Unknown direction: " + direction)
    return zeros


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    # 6492 < x < 6507
    print(solve2('puzzle.txt'))
