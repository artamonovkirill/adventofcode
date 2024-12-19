from d17 import solve, process


def test_solves_sub_examples():
    # bst
    registers = {'B': 0, 'C': 9}
    process(registers, [2, 6])
    assert registers['B'] == 1

    # out
    registers = {'A': 10, 'B': 0, 'C': 0}
    out = process(registers, [5, 0, 5, 1, 5, 4])
    assert out == [0, 1, 2]

    # adv
    registers = {'A': 2024, 'B': 0, 'C': 0}
    out = process(registers, [0, 1, 5, 4, 3, 0])
    assert out == [4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0]
    assert registers['A'] == 0

    registers = {'A': 0, 'B': 29, 'C': 0}
    process(registers, [1, 7])
    assert registers['B'] == 26

    registers = {'A': 0, 'B': 2024, 'C': 43690}
    process(registers, [4, 0])
    assert registers['B'] == 44354

    registers = {'A': 117440, 'B': 0, 'C': 0}
    out = process(registers, [0, 3, 5, 4, 3, 0])
    assert out == [0, 3, 5, 4, 3, 0]


def test_solves_example():
    assert solve('example.txt') == ','.join(str(i) for i in [4, 6, 3, 5, 6, 3, 5, 2, 1, 0])
