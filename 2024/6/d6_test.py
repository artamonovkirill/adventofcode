from d6 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 41
    assert solve2('example.txt') == 6
