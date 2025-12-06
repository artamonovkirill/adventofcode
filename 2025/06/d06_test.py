from d06 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 4277556
    assert solve2('example.txt') == 3263827
