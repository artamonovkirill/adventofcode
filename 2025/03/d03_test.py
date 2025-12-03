from d03 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 357
    assert solve2('example.txt') == 3121910778619
