import static java.lang.Math.abs

class Crabs {
    static List<Integer> best(List<Integer> positions) {
        (0..positions.max()).collect {
            [it, positions.collect {
                p -> abs(p - it)
            }.sum {
                n -> n * (n + 1) / 2
            }]
        }.min { it[1] }
    }

    static void main(String... args) {
        def input = new File('src/main/resources/crabs.txt')
                .text
                .split(',')
                .collect { it as int }
        println(best(input))
    }
}
