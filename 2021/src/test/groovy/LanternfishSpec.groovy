import spock.lang.Specification

class LanternfishSpec extends Specification {
    def 'simulates'() {
        given:
        def input = [3, 4, 3, 1, 2]
        def fish = new Lanternfish(input)

        when:
        fish.simulate(18)

        then:
        fish.population == new Lanternfish(
                [6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8]).population
    }

    def 'simulates a lot'() {
        given:
        def input = [3, 4, 3, 1, 2]
        def fish = new Lanternfish(input)

        when:
        fish.simulate(256)

        then:
        fish.population.values().sum() == 26984457539
    }
}
