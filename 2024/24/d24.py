import re
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()

    initial_values = []
    relations = dict()
    wires = set()

    ops = dict()
    for line in content:
        line = line.rstrip()
        if ':' in line:
            key, value = line.split(': ')
            initial_values.append((key, int(value)))
        elif line:
            matches = re.search('([a-z0-9]{3}) (AND|XOR|OR) ([a-z0-9]{3}) -> ([a-z0-9]{3})', line).groups()
            op = matches[1]
            a = matches[0]
            b = matches[2]
            out = matches[3]

            if a not in relations:
                relations[a] = set()
            relations[a].add(b)

            if b not in relations:
                relations[b] = set()
            relations[b].add(a)

            wires.add((a, op, b, out))

    for a, op, b, out in wires:
        if (a, b) not in ops:
            ops[(a, b)] = dict()
        ops[(a, b)][out] = op

    def set_value(key: str, value: int):
        values[key] = value
        if key in relations:
            for other in relations[key]:
                if other in values:
                    if (key, other) in ops:
                        outs = ops[(key, other)]
                        for out in outs:
                            if out not in values:
                                op = outs[out]
                                if op == 'AND':
                                    out_value = values[key] & values[other]
                                elif op == 'OR':
                                    out_value = values[key] | values[other]
                                elif op == 'XOR':
                                    out_value = values[key] ^ values[other]
                                else:
                                    raise NotImplementedError()
                                set_value(out, out_value)
                    if (other, key) in ops:
                        outs = ops[(other, key)]
                        for out in outs:
                            if out not in values:
                                op = outs[out]
                                if op == 'AND':
                                    out_value = values[other] & values[key]
                                elif op == 'OR':
                                    out_value = values[other] | values[key]
                                elif op == 'XOR':
                                    out_value = values[other] ^ values[key]
                                else:
                                    raise NotImplementedError()
                                set_value(out, out_value)

    values = dict()
    for key, value in initial_values:
        set_value(key, value)

    result = ''
    for i in range(101):
        key = 'z' + str(i).zfill(2)
        if key not in values:
            break
        value = values[key]
        result = str(value) + result
    return int(result, 2)


def solve2(file: str):
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()

    wires = []
    for line in content:
        line = line.rstrip()
        if line and ':' not in line:
            matches = re.search('([a-z0-9]{3}) (AND|XOR|OR) ([a-z0-9]{3}) -> ([a-z0-9]{3})', line).groups()
            op = matches[1]
            a = matches[0]
            b = matches[2]
            out = matches[3]
            wires.append((a, op, b, out))

    certain = {
        142: 105,  # z08
        71: 187,  # z16
        120: 220,  # z28
        116: 72,  # z39
    }

    swaps = []
    for k in certain:
        a1, op1, b1, out1 = wires[k]
        a2, op2, b2, out2 = wires[certain[k]]
        swaps.append(out1)
        swaps.append(out2)
        wires[k] = (a1, op1, b1, out2)
        wires[certain[k]] = (a2, op2, b2, out1)

    outs = dict()
    for a, op, b, out in wires:
        if out in outs:
            raise NotImplementedError()
        outs[out] = (a, op, b)

    def traverse(key: str):
        a, op, b = outs[key]
        if a.startswith('x') and b.startswith('y'):
            return {'op': op, 'left': a, 'right': b}
        if b.startswith('x') and a.startswith('y'):
            return {'op': op, 'left': b, 'right': a}
        t_a = traverse(a)
        t_b = traverse(b)
        if isinstance(t_a['left'], str):
            return {'op': op, 'left': t_a, 'right': t_b}
        return {'op': op, 'left': t_b, 'right': t_a}

    zs = dict()
    for i in range(46):
        key = 'z' + str(i).zfill(2)
        zs[key] = traverse(key)

    for i in range(2, 45):
        index = str(i).zfill(2)
        key = 'z' + index
        assert zs[key]['op'] == 'XOR', key
        assert zs[key]['left']['op'] == 'XOR', key
        assert zs[key]['left']['left'] == 'x' + index, key
        assert zs[key]['left']['right'] == 'y' + index, key

    for z in zs:
        print(z)
        print(to_string(zs[z]) + '\n')

    return ','.join(sorted(swaps))


def to_string(d: dict, indent: int = 0):
    result = ' ' * indent + d['op']
    result += '\n'
    if isinstance(d['left'], str):
        result += ' ' * (indent + 2) + d['left'] + ' ' + d['right']
    else:
        result += to_string(d['left'], indent + 2) + '\n' + to_string(d['right'], indent + 2)
    return result


if __name__ == "__main__":
    print(solve('puzzle.txt'))
    print(solve2('puzzle.txt'))
