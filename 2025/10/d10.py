from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    result = 0
    for line in content:
        parts = line.replace('\n', '').split(' ')
        lamps = tuple(l == '#' for l in parts[0][1:-1])
        switches = []
        for buttons in parts[1:-1]:
            switches.append([int(i) for i in buttons[1:-1].split(',')])
        result += solve_one(lamps, switches)
    return result


def solve_one(lamps: tuple[bool, ...], switches: list[list[int]]):
    current = {lamps}
    i = 0
    while True:
        i += 1
        nxt = set()
        for c in current:
            for s in switches:
                new = list(c)
                for n in s:
                    new[n] = not new[n]
                if all(n == False for n in new):
                    return i
                nxt.add(tuple(new))
        current = nxt

if __name__ == "__main__":
    print(solve('puzzle.txt'))
