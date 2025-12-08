from d08 import solve, solve2


def test_solves_example():
    assert solve('example.txt', 10) == 40
    assert solve2('example.txt') == 25272
