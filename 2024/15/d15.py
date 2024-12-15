from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    walls = set()
    boxes = dict()
    r_x, r_y = 0, 0
    for y, line in enumerate(content):
        line = line.rstrip()
        if line.startswith('#'):
            for x, c in enumerate(line):
                if c == '#':
                    walls.add((x, y))
                elif c == '@':
                    r_x, r_y = x, y
                elif c == 'O':
                    boxes[(x, y)] = True
            width = len(line)
            height = y + 1
        elif line:
            for move in line:
                if move == '>':
                    dx, dy = 1, 0
                elif move == '<':
                    dx, dy = -1, 0
                elif move == '^':
                    dx, dy = 0, -1
                elif move == 'v':
                    dx, dy = 0, 1
                else:
                    raise NotImplementedError()
                next_x, next_y = r_x + dx, r_y + dy
                if (next_x, next_y) not in walls:
                    if (next_x, next_y) not in boxes:
                        r_x, r_y = next_x, next_y
                    else:
                        i = 2
                        c_x, c_y = r_x + dx * i, r_y + dy * i
                        while (c_x, c_y) not in walls:
                            if (c_x, c_y) not in boxes:
                                boxes[(c_x, c_y)] = True
                                del boxes[(next_x, next_y)]
                                r_x, r_y = next_x, next_y
                                break
                            i += 1
                            c_x, c_y = r_x + dx * i, r_y + dy * i

    return sum(x + y * 100 for x, y in boxes)


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    walls = set()
    boxes = dict()
    r_x, r_y = 0, 0

    def symbol(x, y):
        if (x, y) in walls:
            return '#'
        elif (x, y) in boxes:
            return boxes[(x, y)]
        elif x == r_x and y == r_y:
            return '@'
        return '.'

    def find_moves(r_x, r_y, dy):
        moves = set()
        current = {(r_x, r_y + dy)}
        if boxes[(r_x, r_y + dy)] == '[':
            current.add((r_x + 1, r_y + dy))
        elif boxes[(r_x, r_y + dy)] == ']':
            current.add((r_x - 1, r_y + dy))
        while len(current) > 0:
            nxt = set()
            for x, y in current:
                moves.add((x, y, boxes[(x, y)]))
                next_y = y + dy
                if (x, next_y) in walls:
                    return []
                elif (x, next_y) in boxes:
                    nxt.add((x, next_y))
                    if boxes[(x, next_y)] == '[':
                        nxt.add((x + 1, next_y))
                    elif boxes[(x, next_y)] == ']':
                        nxt.add((x - 1, next_y))
            current = nxt
        return moves

    for y, line in enumerate(content):
        line = line.rstrip()
        if line.startswith('#'):
            for x, c in enumerate(line):
                if c == '#':
                    walls.add((2 * x, y))
                    walls.add((2 * x + 1, y))
                elif c == '@':
                    r_x, r_y = 2 * x, y
                elif c == 'O':
                    boxes[(2 * x, y)] = '['
                    boxes[(2 * x + 1, y)] = ']'
            width = 2 * len(line)
            height = y + 1
        elif line:
            print(line)
            for move in line:
                if move == '>':
                    dx, dy = 1, 0
                elif move == '<':
                    dx, dy = -1, 0
                elif move == '^':
                    dx, dy = 0, -1
                elif move == 'v':
                    dx, dy = 0, 1
                else:
                    raise NotImplementedError()
                next_x, next_y = r_x + dx, r_y + dy
                if (next_x, next_y) not in walls:
                    if (next_x, next_y) not in boxes:
                        r_x, r_y = next_x, next_y
                    else:
                        if dy == 0:
                            i = 2
                            c_x = r_x + dx * i
                            while (c_x, r_y) not in walls:
                                if (c_x, r_y) not in boxes:
                                    if dx < 0:
                                        boxes[(c_x, r_y)] = '['
                                    else:
                                        boxes[(c_x, r_y)] = ']'
                                    for j in range(2, i):
                                        p = (r_x + dx * j, r_y)
                                        boxes[p] = '[' if boxes[p] == ']' else ']'
                                    del boxes[(next_x, next_y)]
                                    r_x, r_y = next_x, next_y
                                    break
                                i += 1
                                c_x, c_y = r_x + dx * i, r_y + dy * i
                        else:
                            moves = find_moves(r_x, r_y, dy)
                            if len(moves) > 0:
                                for x, y, _ in moves:
                                    del boxes[(x, y)]
                                for x, y, c in moves:
                                    boxes[(x, y + dy)] = c
                                    r_x, r_y = next_x, next_y

    return sum(x + y * 100 for x, y in boxes if boxes[(x, y)] == '[')


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
