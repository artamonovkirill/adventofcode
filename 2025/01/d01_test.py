from d01 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 3
    assert solve2('example.txt') == 6
