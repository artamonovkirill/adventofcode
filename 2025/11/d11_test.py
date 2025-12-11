from d11 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 5
    assert solve2('example2.txt') == 2
