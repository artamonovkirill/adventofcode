from d16 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 7036
    assert solve2('example.txt') == 45