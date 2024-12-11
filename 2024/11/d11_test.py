from d11 import solve, parse, total


def test_solves_example():
    assert solve('0 1 10 99 999', 1) == parse('1 2024 1 0 9 9 2021976')
    assert solve('125 17', 6) == parse('2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2')
    assert total(solve('125 17', 25)) == 55312
