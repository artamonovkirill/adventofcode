from d13 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 480
    assert [r > 0 for r in solve2('example.txt')] == [False,True,False,True]
