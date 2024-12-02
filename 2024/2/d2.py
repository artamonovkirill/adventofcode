from pathlib import Path


def is_safe(data: list[str]) -> bool:
    first, second, *tail = data
    diff = int(first) - int(second)
    last = int(second)
    if 1 <= abs(diff) <= 3:
        for e in tail:
            new_last = int(e)
            new_diff = last - new_last
            if not (1 <= abs(new_diff) <= 3):
                return False
            if new_diff * diff < 0:
                return False
            last = new_last
            diff = new_diff
        return True
    else:
        return False


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    safes = 0
    for line in content:
        if is_safe(line.split(' ')):
            safes += 1
    return safes


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    safes = 0
    for line in content:
        parts = line.split(' ')
        if is_safe(parts):
            safes += 1
        else:
            for i in range(0, len(parts)):
                if is_safe(parts[:i] + parts[i + 1:]):
                    safes += 1
                    break
    return safes


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
