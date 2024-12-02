from d1 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 11
    assert solve2('example.txt') == 31