import re
from collections import defaultdict
from pathlib import Path


def solve(file: str) -> int:
    path = Path(__file__).parent / file
    with path.open() as f:
        content = f.readlines()
    initial_values = []
    relations = defaultdict(lambda: set())
    processors = defaultdict(lambda: dict())
    values = dict()
    for line in content:
        line = line.rstrip()
        if ':' in line:
            id, value = line.split(': ')
            initial_values.append((id, int(value)))
        elif line:
            matches = re.search('([a-z0-9]{3}) (AND|XOR|OR) ([a-z0-9]{3}) -> ([a-z0-9]{3})', line).groups()
            op = matches[1]
            a = matches[0]
            b = matches[2]
            out = matches[3]
            if op == 'AND':
                fn = ('AND', a, b)
            elif op == 'OR':
                fn = ('OR', a, b)
            elif op == 'XOR':
                fn = ('XOR', a, b)
            else:
                raise NotImplementedError(op)
            relations[a].add(b)
            relations[b].add(a)
            processors[(a, b)][out] = fn
            processors[(b, a)][out] = fn

    def set_value(id, op: tuple):
        if op[0] is None:
            values[id] = op[1]
        elif op[0] == 'AND':
            values[id] = values[op[1]] & values[op[2]]
        elif op[0] == 'OR':
            values[id] = values[op[1]] | values[op[2]]
        elif op[0] == 'XOR':
            values[id] = values[op[1]] ^ values[op[2]]
        else:
            raise NotImplementedError()
        for other in relations[id]:
            if other in values:
                ps = processors[(id, other)]
                for out in ps:
                    if out not in values:
                        set_value(out, ps[out])

    for id, v in initial_values:
        set_value(id, (None, v))

    result = ''
    for i in range(101):
        id = 'z' + str(i).zfill(2)
        if id not in values:
            break
        value = values[id]
        result = str(value) + result
    return int(result, 2)


if __name__ == "__main__":
    print(solve('puzzle.txt'))
