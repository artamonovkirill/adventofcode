from d19 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 6
    assert solve2('example.txt') == 16
