class Segments {
    static List<Integer> decode(List<String> displays) {
        displays.collect { display ->
            def io = display.split('[|]').collect { it.trim().split(' ') }

            def signals = io[0]
            def one = signals.find { it.size() == 2 }
            def four = signals.find { it.size() == 4 }
            if (!one || !four) throw new RuntimeException("Need one and four to decode")

            def output = io[1]
            output.collect { digit -> parse(digit, one, four) }.join('') as int
        }
    }

    private static String parse(String digit, String one, String four) {
        switch (digit.size()) {
            case 2: return 1
            case 4: return 4
            case 3: return 7
            case 7: return 8
            case 5:
                return contains(digit, one)
                        ? 3
                        : contains(digit, four.toList() - one.toList()) ? 5 : 2
            case 6:
                return contains(digit, one)
                        ? (contains(digit, four) ? 9 : 0)
                        : 6
            default: digit
        }
    }

    private static boolean contains(String a, String b) {
        contains(a, b.toList())
    }

    private static boolean contains(String a, List<String> b) {
        a.toList().containsAll(b)
    }

    static void main(String... args) {
        def input = 'src/main/resources/segments.txt' as File
        def readings = input.readLines()
        println(decode(readings).sum())
    }

}
