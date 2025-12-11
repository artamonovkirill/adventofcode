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


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    links = defaultdict(lambda: [])
    for line in content:
        frm, tos = line.replace('\n', '').split(': ')
        for to in tos.split(' '):
            links[frm].append(to)

    fft_dac = paths(links, 'svr', 'fft') * paths(links, 'fft', 'dac') * paths(links, 'dac', 'out')
    dac_fft = paths(links, 'svr', 'dac') * paths(links, 'dac', 'fft') * paths(links, 'fft', 'out')
    return fft_dac + dac_fft


def paths(links: dict[str, list[str]], start: str, end: str):
    result = 0
    current = {start: 1}
    while current:
        nxt = dict()
        for frm, count in current.items():
            for to in links[frm]:
                if to not in nxt:
                    nxt[to] = count
                else:
                    nxt[to] += count
        if end in nxt:
            result += nxt[end]
        current = nxt
    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
