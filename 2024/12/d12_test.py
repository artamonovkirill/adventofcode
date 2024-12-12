from d12 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 140
    assert solve2('example.txt') == 80
    assert solve2('example2.txt') == 236
    assert solve2('example3.txt') == 368
