import sys
from collections import defaultdict
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    points = []
    for line in content:
        x, y = line.replace('\n', '').split(',')
        points.append((int(x), int(y)))
    result = 0
    for i in range(len(points)):
        for j in range(i + 1, len(points)):
            x1, y1 = points[i]
            x2, y2 = points[j]
            area = (abs(x1 - x2) + 1) * (abs(y1 - y2) + 1)
            if area > result:
                result = area
    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
