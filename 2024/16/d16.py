import sys
from pathlib import Path

sys.path.append(str(Path(__file__).parent.parent.parent))


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    walls = set()
    start_x, start_y = None, None
    end = None
    for y, line in enumerate(content):
        for x, c in enumerate(line.rstrip()):
            if c == '#':
                walls.add((x, y))
            elif c == 'S':
                start_x, start_y = x, y
            elif c == 'E':
                end = x, y

    scores = dict()
    scores[(start_x, start_y)] = 0
    current = {(start_x, start_y): (1, 0)}
    while len(current) > 0:
        nxt = dict()
        for x, y in current:
            # go straight
            dx, dy = current[(x, y)]
            next_x = x + dx
            next_y = y + dy
            if (next_x, next_y) not in walls:
                if (next_x, next_y) not in scores or scores[(next_x, next_y)] > scores[(x, y)] + 1:
                    nxt[(next_x, next_y)] = (dx, dy)
                    scores[(next_x, next_y)] = scores[(x, y)] + 1

            # turn left
            if dx == 0:
                if dy == 1:  # south
                    next_dx, next_dy = 1, 0
                else:  # north
                    next_dx, next_dy = -1, 0
            else:
                if dx == 1:  # east
                    next_dx, next_dy = 0, -1
                else:  # west
                    next_dx, next_dy = 0, 1
            next_x = x + next_dx
            next_y = y + next_dy
            if (next_x, next_y) not in walls:
                if (next_x, next_y) not in scores or scores[(next_x, next_y)] > scores[(x, y)] + 1001:
                    nxt[(next_x, next_y)] = (next_dx, next_dy)
                    scores[(next_x, next_y)] = scores[(x, y)] + 1001

            # turn right
            if dx == 0:
                if dy == 1:  # south
                    next_dx, next_dy = -1, 0
                else:  # north
                    next_dx, next_dy = 1, 0
            else:
                if dx == 1:  # east
                    next_dx, next_dy = 0, 1
                else:  # west
                    next_dx, next_dy = 0, -1
            next_x = x + next_dx
            next_y = y + next_dy
            if (next_x, next_y) not in walls:
                if (next_x, next_y) not in scores or scores[(next_x, next_y)] > scores[(x, y)] + 1001:
                    nxt[(next_x, next_y)] = (next_dx, next_dy)
                    scores[(next_x, next_y)] = scores[(x, y)] + 1001
        current = nxt

    return scores[end]


if __name__ == "__main__":
    print(solve('puzzle.txt'))
