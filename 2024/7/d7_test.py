from d7 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 3749
    assert solve2('example.txt') == 11387
