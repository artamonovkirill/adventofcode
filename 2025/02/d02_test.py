from d02 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 1227775554
    assert solve2('example.txt') == 4174379265
