import static java.lang.Math.floorDiv

class ALU {
    static final INPUT = 'src/main/resources/alu.txt' as File
    static final INSTRUCTIONS = INPUT.text.split(/inp /).findAll { it }

    static long check(long n) {
        assert INSTRUCTIONS.size() == 14
        def digits = digits(n)
        def init = ['x', 'y', 'z', 'w'].collectEntries { [(it): 0 as long] }
        digits.indexed().inject(init) { acc, i, digit ->
            acc += ['w': digit]
            def commands = instruction(i)
            apply(acc, commands)
        }['z']
    }

    static List<Integer> digits(long n) {
        n.toString().toList().collect { it as int }
    }

    static List<String> instruction(int n) {
        INSTRUCTIONS[n].split('\n').toList().drop(1)
    }

    static Map<String, Long> apply(Map<String, Long> vars, List<String> instructions) {
        vars = vars.collectEntries { it }
        instructions.each {
            def (action, a, b) = it.split(' ').toList()
            def f = parse(action)
            assert a =~ /[wxyz]/
            vars[a] = f(vars[a], b =~ /-?[0-9]+/ ? b as long : vars[b])
        }
        assert vars.size() == 4
        return vars.asImmutable()
    }

    static void main(String... args) {
        (11..90).each {
            def (fst, snd) = digits(it)
            def valid = generate([0: ''], fst, snd, 0)
            if (valid)
                throw new RuntimeException("Found: ${valid.values().min()}")
            else
                println("No valid numbers found")
        }
    }

    static Map<Integer, String> generate(Map<Integer, String> acc, int fst, int snd, int step = 0) {
        if (step > 13) {
            println("Max: ${fst}${snd}${'9'.repeat(12)}")
            return acc.findAll {
                k, v -> k == 0
            }
        }
        def result = [:]
        def commands = instruction(step)
        def range = step == 0 ? [fst] : step == 1 ? [snd] : (1..9)
        range.each { d ->
            acc.each { k, v ->
                def vars = [x: 0, y: 0, z: k, w: d]
                def c = apply(vars, commands)['z']
                if (!result.containsKey(c) || result[c] > v + d)
                    result[c] = v + d
            }
        }
        return generate(result, fst, snd, step + 1)
    }

    static Closure<Long> parse(String action) {
        switch (action) {
            case 'add': return { a, b -> a + b }
            case 'mul': return { a, b -> a * b }
            case 'div': return { long a, long b -> floorDiv(a, b) }
            case 'mod': return { a, b -> a % b }
            case 'eql': return { a, b -> a == b ? 1 : 0 }
            default: throw new RuntimeException("not implemented for '$action'")
        }
    }

}
