from d10 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 36
    assert solve2('example.txt') == 81
