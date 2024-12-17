import re
from pathlib import Path


class Halted(Exception):
    pass


def process_one(registers: dict[str, int], program: list[int], out: list[int], i):
    if i < 0 or i >= len(program):
        raise Halted()
    opcode = program[i]
    operand = program[i + 1]

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


def process(registers: dict[str, int], program: list[int], i=0) -> list[int]:
    out = []
    while True:
        try:
            i = process_one(registers, program, out, i)
        except Halted:
            return out


def solve(file: str) -> str:
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
    out = process(registers, program)
    return ','.join(str(o) for o in out)


def solve2(file: str) -> int:
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

    not_the_solution = set()

    a = 0
    while True:
        rs = {'A': a, 'B': registers['B'], 'C': registers['C']}
        start_key = rs['A'], rs['B'], rs['C']
        if start_key not in not_the_solution:
            i = 0
            out = []
            while True:
                before_key = start_key
                if before_key in not_the_solution:
                    a += 1
                    break
                try:
                    i = process_one(rs, program, out, i)
                except Halted:
                    if out == program:
                        return a
                    else:
                        not_the_solution.add(before_key)
                        a += 1
                        break
                if len(out) > 0:
                    if len(out) > len(program) or out[len(out) - 1] != program[len(out) - 1]:
                        not_the_solution.add(before_key)
                        a += 1
                        break


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    # > 52 000 000 000
    print(solve2('puzzle.txt'))
