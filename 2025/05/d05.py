from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        fresh, ingredients = f.read().split('\n\n')
    ranges = []
    for line in fresh.split('\n'):
        left, right = line.split('-')
        ranges.append((int(left), int(right)))
    fresh = 0
    for ingredient in ingredients.split('\n'):
        i = int(ingredient)
        for left, right in ranges:
            if left <= i <= right:
                fresh += 1
                break
    return fresh


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        fresh, _ = f.read().split('\n\n')
    intervals = dict()
    candidates = []
    for line in fresh.split('\n'):
        left, right = line.split('-')
        candidates.append((int(left), int(right)))
    for left, right in sorted(candidates):
        while True:
            for existing in intervals:
                existing_left, existing_right = existing
                if left <= existing_left <= right <= existing_right:
                    del intervals[(existing_left, existing_right)]
                    right = existing_right
                    break
                elif existing_left <= left <= existing_right <= right:
                    del intervals[(existing_left, existing_right)]
                    left = existing_left
                    break
                elif existing_left <= left <= right <= existing_right:
                    del intervals[(existing_left, existing_right)]
                    left = existing_left
                    right = existing_right
                    break
                elif left <= existing_left <= existing_right <= right:
                    del intervals[(existing_left, existing_right)]
                    break
            intervals[(left, right)] = True
            break
    result = 0
    for interval in intervals:
        left, right = interval
        result += right - left + 1
    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
