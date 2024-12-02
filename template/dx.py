from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    for line in content:
        print(line)
    return 0


if __name__ == "__main__":
    print(solve('puzzle.txt'))
