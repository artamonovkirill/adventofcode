class Segments {
    static List<Integer> decode(List<String> data) {
        data.collect { line ->
            def io = line.split('[|]')
            def input = io[0].trim().split(' ')
            def output = io[1].trim().split(' ')
            def all = input + output
            def one = all.find { it.size() == 2 }
            def four = all.find { it.size() == 4 }
            if (!one || !four) throw new RuntimeException("Need one and four to decode")
            output.collect {
                d -> parse(d, one, four)
            }.join('') as int
        }
    }

    private static String parse(String d, String one, String four) {
        if (d.size() == 2) return 1
        if (d.size() == 4) return 4
        if (d.size() == 3) return 7
        if (d.size() == 7) return 8
        if (d.size() == 5) {
            return d.toList().containsAll(one.toList())
                    ? 3
                    : d.toList().containsAll(four.toList() - one.toList()) ? 5 : 2

        }
        if (d.size() == 6) {
            return d.toList().containsAll(one.toList())
                    ? (d.toList().containsAll(four.toList()) ? 9 : 0)
                    : 6
        }
        return d
    }

    static void main(String... args) {
        def input = 'src/main/resources/segments.txt' as File
        def readings = input.readLines()
        println(decode(readings).sum())
    }

}
