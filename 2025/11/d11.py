from collections import defaultdict
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    links = defaultdict(lambda: [])
    for line in content:
        frm, tos = line.replace('\n', '').split(': ')
        for to in tos.split(' '):
            links[frm].append(to)
    result = 0
    current = ['you']
    while current:
        nxt = []
        for frm in current:
            for to in links[frm]:
                if to == 'out':
                    result += 1
                else:
                    nxt.append(to)
        current = nxt
    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
