class Lanternfish {
    Map<Integer, Long> population

    Lanternfish(List<Integer> input) {
        population = input.groupBy { it }
                .collectEntries { [(it.key): it.value.size() as long] }
    }

    void simulate(int days) {
        (1..days).each {
            population = population.collectMany {
                it.key == 0
                        ? [new Pair(a: 6, b: it.value), new Pair(a: 8, b: it.value)]
                        : [new Pair(a: it.key - 1, b: it.value)]
            }.groupBy { it.a }.collectEntries {
                [(it.key): it.value.sum { it.b }]
            }
        }
    }

    static void main(String... args) {
        def input = [
                5, 1, 1, 4, 1, 1, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 2, 1, 1, 1,
                3, 5, 1, 1, 1, 5, 4, 1, 1, 1, 2, 2, 1, 1, 1, 2, 1, 1, 1, 2, 5, 2, 1,
                2, 2, 3, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 1, 4, 1, 1, 1, 5, 4, 1, 1, 3,
                3, 2, 1, 1, 1, 5, 1, 1, 4, 1, 1, 5, 1, 1, 5, 1, 2, 3, 1, 5, 1, 3, 2,
                1, 3, 1, 1, 4, 1, 1, 1, 1, 2, 1, 2, 1, 1, 2, 1, 1, 1, 4, 4, 1, 5, 1,
                1, 3, 5, 1, 1, 5, 1, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 3, 1, 1, 1,
                1, 1, 2, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 5, 1, 1, 1, 1, 4, 1, 1, 1, 1,
                4, 1, 1, 1, 1, 3, 1, 2, 1, 2, 1, 3, 1, 3, 4, 1, 1, 1, 1, 1, 1, 1, 5,
                1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 1, 2, 2, 1, 2, 4, 1, 1, 3, 1, 1, 1, 5,
                1, 3, 1, 1, 1, 5, 5, 1, 1, 1, 1, 2, 3, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1,
                1, 1, 1, 5, 1, 4, 3, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1,
                1, 1, 1, 1, 1, 1, 1, 1, 3, 3, 1, 2, 2, 1, 4, 1, 5, 1, 5, 1, 1, 1, 1,
                1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 1, 1, 4, 3, 1, 1, 4]
        def fish = new Lanternfish(input)
        fish.simulate(256)
        println(fish.population.values().sum())
    }
}
