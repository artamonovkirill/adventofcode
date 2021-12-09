import spock.lang.Specification

class SmokeCaveSpec extends Specification {
    def 'solves example'() {
        given:
        def cave = new SmokeCave([
                '2199943210',
                '3987894921',
                '9856789892',
                '8767896789',
                '9899965678'])

        expect:
        cave.lowPoints() == [[1, 1, 0], [0, 9, 0], [5, 2, 2], [5, 6, 4]]
        cave.risk() == 15
        cave.basins()*.size() == [3, 9, 14, 9]
        cave.basinsScore() == 1134
    }
}
