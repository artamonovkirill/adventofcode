import sys
import textwrap
from collections import Counter, defaultdict
from pathlib import Path

sys.path.append(str(Path(__file__).parent.parent.parent))


def parse(data: str) -> dict[int, int]:
    return dict(Counter(int(d) for d in data.split(' ')).items())


def solve(data: str, n: int) -> dict[int, int]:
    current = parse(data)
    for i in range(n):
        nxt = defaultdict(lambda: 0)
        for c in current:
            if c == 0:
                nxt[1] += current[c]
            else:
                s = str(c)
                if len(s) % 2 == 0:
                    for w in textwrap.wrap(s, len(s) // 2):
                        nxt[int(w)] += current[c]
                else:
                    nxt[c * 2024] += current[c]
        current = nxt
    return current


def total(d: dict[int, int]) -> int:
    return sum(d[k] for k in d)


if __name__ == "__main__":
    print(total(solve('3279 998884 1832781 517 8 18864 28 0', 25)))
    print(total(solve('3279 998884 1832781 517 8 18864 28 0', 75)))
