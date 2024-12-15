from d15 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 10092
    assert solve2('example.txt') == 9021
