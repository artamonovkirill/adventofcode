import re
from pathlib import Path


class Halted(Exception):
    pass


def process_one(registers: dict[str, int], program: list[int], out: list[int], i):
    if i < 0 or i >= len(program):
        raise Halted()
    opcode = program[i]
    operand = program[i + 1]

    value = parse_operand(operand, registers)

    if opcode == 0:  # adv
        numerator = registers['A']
        denominator = pow(2, value)
        registers['A'] = numerator // denominator
        i += 2
    elif opcode == 1:  # bxl
        registers['B'] = registers['B'] ^ operand
        i += 2
    elif opcode == 2:  # bst
        registers['B'] = value % 8
        i += 2
    elif opcode == 3:  # jnz
        if registers['A'] != 0:
            i = operand
        else:
            i += 2
    elif opcode == 4:  # bxc
        registers['B'] = registers['B'] ^ registers['C']
        i += 2
    elif opcode == 5:  # out
        out.append(value % 8)
        i += 2
    elif opcode == 6:  # bdv
        numerator = registers['A']
        denominator = pow(2, value)
        registers['B'] = numerator // denominator
        i += 2
    elif opcode == 7:  # cdv
        numerator = registers['A']
        denominator = pow(2, value)
        registers['C'] = numerator // denominator
        i += 2
    else:
        raise NotImplementedError('opcode', opcode)
    return i


def parse_operand(operand, registers):
    if operand == 0:
        value = 0
    elif operand == 1:
        value = 1
    elif operand == 2:
        value = 2
    elif operand == 3:
        value = 3
    elif operand == 4:
        value = registers['A']
    elif operand == 5:
        value = registers['B']
    elif operand == 6:
        value = registers['C']
    elif operand == 7:
        value = None
    else:
        raise NotImplementedError('operand', operand)
    return value


def process(registers: dict[str, int], program: list[int]) -> list[int]:
    i = 0
    out = []
    while True:
        try:
            i = process_one(registers, program, out, i)
        except Halted:
            return out


def solve(file: str) -> str:
    program, registers = parse(file)
    out = process(registers, program)
    return ','.join(str(o) for o in out)


def parse(file):
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    registers = dict()
    for line in content:
        line = line.rstrip()
        if line.startswith('Register'):
            match = re.search(r'Register ([A-C]): ([0-9]+)', line)
            registers[match.group(1)] = int(match.group(2))
        elif line.startswith('Program'):
            match = re.search(r'Program: (.*)', line)
            program = [int(c) for c in match.group(1).split(',')]
    assert program
    return program, registers


def solve2(file: str) -> int:
    program = parse(file)[0]

    current = {0}
    expected = list(reversed(program))
    i = 0

    while i < len(expected):
        nexts = set()
        target_out = expected[i]
        for target_a in current:
            for j in range(8):
                a = j + 8 * target_a
                print(a, target_a, target_out, process_once(a, program))
                if process_once(a, program) == (target_a, target_out):
                    nexts.add(a)
        current = nexts
        i += 1

    return min(current)


def process_once(a: int, program: list[int]):
    i = 0
    out = []
    registers = {'A': a}

    while 0 <= i < len(program):
        opcode = program[i]
        operand = program[i + 1]
        value = parse_operand(operand, registers)

        if opcode == 0:  # adv
            numerator = registers['A']
            denominator = pow(2, value)
            registers['A'] = numerator // denominator
            i += 2
        elif opcode == 1:  # bxl
            registers['B'] = registers['B'] ^ operand
            if registers['B'] > 7 or registers['B'] < 0:
                raise NotImplementedError()
            i += 2
        elif opcode == 2:  # bst
            registers['B'] = value % 8
            i += 2
        elif opcode == 3:  # jnz
            if len(out) > 1:
                raise NotImplementedError()
            return registers['A'], out[0]
        elif opcode == 4:  # bxc
            registers['B'] = registers['B'] ^ registers['C']
            i += 2
        elif opcode == 5:  # out
            out.append(value % 8)
            i += 2
        elif opcode == 6:  # bdv
            numerator = registers['A']
            denominator = pow(2, value)
            registers['B'] = numerator // denominator
            i += 2
        elif opcode == 7:  # cdv
            numerator = registers['A']
            denominator = pow(2, value)
            registers['C'] = numerator // denominator
            i += 2
        else:
            raise NotImplementedError('opcode', opcode)
    raise Halted()


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
