import spock.lang.Specification

class ChitonSpec extends Specification {
    def 'solves example'() {
        given:
        def input = 'src/test/resources/chiton/1.txt' as File
        def solution = new Chiton(input)

        expect:
        solution.best() == 40
    }

    def 'extends example'() {
        given:
        def input = 'src/test/resources/chiton/1.txt' as File
        def solution = new Chiton(input)
        and:
        def oracle = 'src/test/resources/chiton/extended.txt' as File

        when:
        def extended = solution.extend()

        then:
        extended.cave == new Chiton(oracle).cave
        extended.best() == 315
    }

    def 'solves another example'() {
        given:
        def input = 'src/test/resources/chiton/2.txt' as File
        def solution = new Chiton(input)

        expect:
        solution.best() == 8
    }
}
