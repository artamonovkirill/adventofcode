from d04 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 13
    assert solve2('example.txt') == 43
