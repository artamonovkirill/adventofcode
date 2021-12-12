import spock.lang.Specification

class PassageSpec extends Specification {
    def 'finds paths in very simple example'() {
        given:
        def connections = [
                'start-A',
                'start-b',
                'A-c',
                'A-b',
                'b-d',
                'A-end',
                'b-end']
        and:
        def oracle = 'src/test/resources/passage/simple.txt' as File
        def expected = oracle.readLines() as Set

        expect:
        new Passage(connections).paths() == expected.toSet()
    }

    def 'finds paths in a medium example'() {
        given:
        def connections = [
                'dc-end',
                'HN-start',
                'start-kj',
                'dc-start',
                'dc-HN',
                'LN-dc',
                'HN-end',
                'kj-sa',
                'kj-HN',
                'kj-dc']

        expect:
        new Passage(connections).paths().size() == 103
    }

    def 'finds paths in a complex example'() {
        given:
        def input = 'src/test/resources/passage/complex.txt' as File
        def connections = input.readLines()

        expect:
        new Passage(connections).paths().size() == 3509
    }
}
