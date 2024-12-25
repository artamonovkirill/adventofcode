from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    lines = []
    keys = []
    locks = []
    for line in content:
        line = line.rstrip()
        if line:
            lines.append(line)
            if len(lines) == 7:
                if lines[0] == '#####':
                    lock = []
                    for x in range(5):
                        for y in range(1, 7):
                            if lines[y][x] == '.':
                                lock.append(y - 1)
                                break
                    locks.append(lock)
                elif lines[6] == '#####':
                    key = []
                    for x in range(5):
                        for y in range(1, 7):
                            if lines[6 - y][x] == '.':
                                key.append(y - 1)
                                break
                    keys.append(key)
                else:
                    raise NotImplementedError()
                lines = []

    count = 0
    for lock in locks:
        for key in keys:
            if fit(key, lock):
                count += 1

    return count


def fit(key, lock):
    for x in range(5):
        if lock[x] + key[x] > 5:
            return False
    return True


if __name__ == "__main__":
    print(solve('puzzle.txt'))
