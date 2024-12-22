from pathlib import Path


def process(secret: int) -> int:
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
        for _ in range(2000):
            secret = process(secret)
        total += secret
    return total


def prices(secret: int, n: int) -> list[int]:
    ps = []
    for _ in range(n):
        ps.append(secret % 10)
        secret = process(secret)
    return ps


def solve2(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    all_sequences = set()
    all_bananas = []
    for line in content:
        secret = int(line.rstrip())
        ps = prices(secret, 2000)
        seq = (ps[1] - ps[0], ps[2] - ps[1], ps[3] - ps[2], ps[4] - ps[3])
        all_sequences.add(seq)
        bananas = {seq: ps[4]}
        for i in range(5, 2000):
            seq = (seq[1], seq[2], seq[3], ps[i] - ps[i - 1])
            if seq not in bananas:
                all_sequences.add(seq)
                bananas[seq] = ps[i]
        all_bananas.append(bananas)

    best = None
    for seq in all_sequences:
        total = sum(b[seq] if seq in b else 0 for b in all_bananas)
        if best is None or best < total:
            best = total
    return best


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
