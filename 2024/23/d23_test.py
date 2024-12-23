from d23 import solve, solve2


def test_solves_example():
    assert solve('example.txt') == 7
    assert solve2('example.txt') == 'co,de,ka,ta'
