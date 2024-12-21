from d21 import type_on, type_sequentially, solve, numeric_keyboard, \
    directional_keyboard


def test_solves_example():
    expected = {'<A^A>^^AvvvA', '<A^A^>^AvvvA', '<A^A^^>AvvvA'}
    assert type_on(numeric_keyboard, '029A') == (expected, 12)
    assert type_sequentially('029A', [numeric_keyboard]) == (expected, 12)

    two_layers = [numeric_keyboard, directional_keyboard]
    paths, size = type_sequentially('029A', two_layers)
    assert 'v<<A>>^A<A>AvA<^AA>A<vAAA>^A' in paths
    assert size == 28

    three_layers = [numeric_keyboard, directional_keyboard, directional_keyboard]

    paths, size = type_sequentially('029A', three_layers)
    assert '<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A' in paths
    assert size == 68

    paths, size = type_sequentially('980A', three_layers)
    assert '<v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A' in paths

    assert solve('example.txt') == 126384
