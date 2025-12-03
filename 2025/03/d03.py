from pathlib import Path


def solve_line(line: str) -> int:
    max_first = 0
    max_first_i = 0
    for i in range(0, len(line) - 1):
        value = int(line[i])
        if value > max_first:
            max_first = value
            max_first_i = i
    max_second = 0
    for i in range(max_first_i + 1, len(line)):
        value = int(line[i])
        if value > max_second:
            max_second = value

    return max_first * 10 + max_second


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    result = 0
    for line in content:
        result += solve_line(line.replace('\n', ''))
    return result

def solve_line2(line: str) -> int:
    result = 0
    left_bound = 0
    for tail_len in reversed(range(0, 12)):
        current_max = 0
        for i in range(left_bound, len(line) - tail_len):
            value = int(line[i])
            if value > current_max:
                left_bound = i + 1
                current_max = value
        result = result*10 + current_max
    return result

def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    result = 0
    for line in content:
        result += solve_line2(line.replace('\n', ''))
    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
