import spock.lang.Specification

class OctopusesSpec extends Specification {
    def 'solves simple example'() {
        given:
        def input = 'src/test/resources/octopuses/simple/0.txt' as File
        def solution = new Octopuses(input)
        and:
        def oracle = "src/test/resources/octopuses/simple/${steps}.txt" as File
        def expected = new Octopuses(oracle)

        when:
        def flashes = solution.advance(steps)

        then:
        solution.octopuses == expected.octopuses
        flashes == expectedFlashes

        where:
        steps | expectedFlashes
        1     | 9
        2     | 9
    }

    def 'solves example'() {
        given:
        def input = 'src/test/resources/octopus.txt' as File
        def solution = new Octopuses(input)
        and:
        def oracle = "src/test/resources/octopuses/${steps}.txt" as File
        def expected = new Octopuses(oracle)

        when:
        def flashes = solution.advance(steps)

        then:
        solution.octopuses == expected.octopuses
        flashes == expectedFlashes

        where:
        steps | expectedFlashes
        1 | 0
        2 | 35
        3 | 80
        4 | 96
    }

    def 'solves 100 steps'() {
        given:
        def input = 'src/test/resources/octopus.txt' as File
        def solution = new Octopuses(input)

        expect:
        solution.advance(100) == 1656
    }

    def 'finds simultaneous flash'() {
        given:
        def input = 'src/test/resources/octopus.txt' as File
        def solution = new Octopuses(input)

        expect:
        solution.all() == 195
    }
}
