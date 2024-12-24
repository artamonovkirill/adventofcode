from d24 import solve


def test_solves_example():
    assert solve('example.txt') == 4
    assert solve('example2.txt') == 2024
