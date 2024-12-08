from d8 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 2
    assert solve('example2.txt') == 14
    assert solve2('example3.txt') == 9
    assert solve2('example2.txt') == 34
