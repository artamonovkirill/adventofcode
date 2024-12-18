from d18 import solve, solve2


def test_solves_example():
    assert solve('example.txt', 12, 7) == 22
    assert solve2('example.txt', 12, 7) == (6, 1)
