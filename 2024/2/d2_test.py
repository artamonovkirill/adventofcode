from d2 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 2
    assert solve2('example.txt') == 4
