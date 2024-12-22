from d22 import solve, prices, solve2


def test_solves_example():
    assert solve('example.txt') == 37327623
    assert solve2('example.txt') == 24
