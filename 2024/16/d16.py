from pathlib import Path

from util.d2 import neighbours


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

    scores = {(start_x, start_y): 0}
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


def solve2(file: str) -> int:
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

    height = len(content)
    width = len(content[0].rstrip())

    crossings = set()
    for y in range(height):
        for x in range(width):
            if (x, y) not in walls:
                vertical, horizontal = 0, 0
                for next_dx, dy in neighbours():
                    next_x = x + next_dx
                    next_y = y + dy
                    if (next_x, next_y) not in walls:
                        if dy == 0:
                            horizontal += 1
                        else:
                            vertical += 1
                assert vertical + horizontal <= 4
                if vertical + horizontal > 2:
                    crossings.add((x, y))

    scores = {(start_x, start_y): 0}
    current = {(start_x, start_y): (1, 0)}
    while len(current) > 0:
        nxts = dict()
        for x, y in current:
            dx, dy = current[(x, y)]
            for next_dx, next_dy in neighbours():
                next_x, next_y = x + next_dx, y + next_dy
                if (next_x, next_y) not in walls:
                    if (next_x, next_y) not in crossings:
                        next_score = scores[(x, y)] + 1
                        if (dx, dy) != (next_dx, next_dy):
                            next_score += 1000
                        if (next_x, next_y) not in scores or next_score < scores[(next_x, next_y)]:
                            nxts[(next_x, next_y)] = (next_dx, next_dy)
                            scores[(next_x, next_y)] = next_score
                    else:
                        for next_next_dx, next_next_dy in neighbours():
                            next_next_x, next_next_y = next_x + next_next_dx, next_y + next_next_dy
                            if (next_next_x, next_next_y) not in walls:
                                next_next_score = scores[(x, y)] + 2
                                if (dx, dy) != (next_dx, next_dy) or (next_dx, next_dy) != (next_next_dx, next_next_dy):
                                    next_next_score += 1000
                                if (next_next_x, next_next_y) not in scores or next_next_score < scores[
                                    (next_next_x, next_next_y)]:
                                    nxts[(next_next_x, next_next_y)] = (next_next_dx, next_next_dy)
                                    scores[(next_next_x, next_next_y)] = next_next_score
        current = nxts

    seats = {end}
    current = {end}
    while len(current) > 0:
        nxts = set()
        for x, y in current:
            for dx, dy in neighbours():
                next_x, next_y = x + dx, y + dy
                if (next_x, next_y) in scores:
                    if scores[(next_x, next_y)] == scores[(x, y)] - 1:
                        seats.add((next_x, next_y))
                        nxts.add((next_x, next_y))
                    if scores[(next_x, next_y)] == scores[(x, y)] - 1001:
                        seats.add((next_x, next_y))
                        nxts.add((next_x, next_y))
                elif (next_x, next_y) in crossings:
                    for next_next_dx, next_next_dy in neighbours():
                        next_next_x, next_next_y = next_x + next_next_dx, next_y + next_next_dy
                        if (next_next_x, next_next_y) in scores:
                            if scores[(next_next_x, next_next_y)] == scores[(x, y)] - 2:
                                seats.add((next_next_x, next_next_y))
                                nxts.add((next_next_x, next_next_y))
                            if scores[(next_next_x, next_next_y)] == scores[(x, y)] - 1002:
                                seats.add((next_next_x, next_next_y))
                                nxts.add((next_next_x, next_next_y))

        current = nxts

    for (x, y) in crossings:
        for dx, dy in neighbours():
            next_x, next_y = x + dx, y + dy
            if (next_x, next_y) in seats:
                seats.add((x, y))
                break

    return len(seats)


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
