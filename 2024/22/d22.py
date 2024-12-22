from pathlib import Path


def generate(secret: int, n: int) -> int:
    for _ in range(n):
        secret = secret ^ (secret * 64)
        secret %= 16777216
        secret = secret ^ (secret // 32)
        secret %= 16777216
        secret = secret ^ (secret * 2048)
        secret %= 16777216
    return secret


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    total = 0
    for line in content:
        secret = int(line.rstrip())
        total += generate(secret, 2000)
    return total


if __name__ == "__main__":
    print(solve('puzzle.txt'))
