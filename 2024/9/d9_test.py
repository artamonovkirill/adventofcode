from d9 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 1928
    assert solve2('example.txt') == 2858
