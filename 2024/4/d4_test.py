from d4 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 4
    assert solve('example2.txt') == 18
    assert solve('example3.txt') == 12
    assert solve('example4.txt') == 12
    assert solve2('example2.txt') == 9
