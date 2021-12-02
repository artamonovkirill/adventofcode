import spock.lang.Specification

class DiveSpec extends Specification {
    def 'calculates position'() {
        when:
        def position = Dive.execute([
                'forward 5',
                'down 5',
                'forward 8',
                'up 3',
                'down 8',
                'forward 2'])

        then:
        position.x == 15
        position.depth == 60
    }
}
