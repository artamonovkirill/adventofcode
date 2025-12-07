from d07 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 21
    assert solve2('example.txt') == 40
