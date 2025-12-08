from collections import defaultdict
from pathlib import Path


def solve(file: str, limit: int) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()

    boxes = []
    for line in content:
        x, y, z = line.replace('\n', '').split(',')
        boxes.append((int(x), int(y), int(z)))

    distances = defaultdict(lambda: [])
    for i in range(len(boxes)):
        for j in range(i + 1, len(boxes)):
            distance = (boxes[i][0] - boxes[j][0]) ** 2 + (boxes[i][1] - boxes[j][1]) ** 2 + (boxes[i][2] - boxes[j][2]) ** 2
            distances[distance].append((boxes[i], boxes[j]))
    for distance, pairs in distances.items():
        assert len(pairs) == 1

    circuits = {box: {box} for box in boxes}

    for distance in sorted(distances.keys())[:limit]:
        frm, to = distances[distance][0]
        joined = circuits[frm] | circuits[to]
        for box in joined:
            circuits[box] = joined

    uniq = dict()
    for circuit in circuits.values():
        uniq[tuple(sorted(circuit))] = len(circuit)

    result = 1
    for length in sorted(uniq.values(), reverse=True)[:3]:
        result *= length
    return result


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()

    boxes = []
    for line in content:
        x, y, z = line.replace('\n', '').split(',')
        boxes.append((int(x), int(y), int(z)))

    distances = defaultdict(lambda: [])
    for i in range(len(boxes)):
        for j in range(i + 1, len(boxes)):
            distance = (boxes[i][0] - boxes[j][0]) ** 2 + (boxes[i][1] - boxes[j][1]) ** 2 + (boxes[i][2] - boxes[j][2]) ** 2
            distances[distance].append((boxes[i], boxes[j]))
    for distance, pairs in distances.items():
        assert len(pairs) == 1

    circuits = {box: {box} for box in boxes}

    for distance in sorted(distances.keys()):
        frm, to = distances[distance][0]
        joined = circuits[frm] | circuits[to]
        if len(joined) == len(boxes):
            return frm[0] * to[0]
        for box in joined:
            circuits[box] = joined

    raise Exception('no solution found')


if __name__ == "__main__":
    print(solve('puzzle.txt', 1000))
    print(solve2('puzzle.txt'))
