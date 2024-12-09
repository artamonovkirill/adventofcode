from collections import defaultdict
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.read()
    i = 0
    data = dict()
    file = True
    id = 0
    for c in content:
        if file:
            for _ in range(int(c)):
                data[i] = id
                i += 1
            id += 1
        else:
            i += int(c)
        file = not file
    left = 0
    right = i
    while left < right:
        if left not in data and right in data:
            data[left] = data[right]
            del data[right]
        if left in data:
            left += 1
        if right not in data:
            right -= 1
    result = 0
    for i in data:
        result += i * data[i]
    return result


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.read()
    pointer = 0
    i_to_id = defaultdict(lambda: dict())
    id_start = dict()
    id_end = dict()
    file = True
    id = 0
    for c in content:
        if file:
            id_start[id] = pointer
            for _ in range(int(c)):
                i_to_id[pointer] = id
                pointer += 1
            id_end[id] = pointer
            id += 1
        else:
            pointer += int(c)
        file = not file
    id -= 1
    unmovable = set()
    while id >= 0:
        left = 0
        start = id_start[id]
        length = id_end[id] - start
        if length in unmovable:
            id -= 1
            continue
        while True:
            while left in i_to_id:
                left += 1
            if left + length - 1 >= start:
                unmovable.add(length)
                break
            if all(left + i not in i_to_id for i in range(length)):
                for i in range(length):
                    i_to_id[left + i] = id
                    del i_to_id[start + i]
                break
            else:
                while left not in i_to_id:
                    left += 1
        id -= 1
    result = 0
    for i in i_to_id:
        result += i * i_to_id[i]
    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
