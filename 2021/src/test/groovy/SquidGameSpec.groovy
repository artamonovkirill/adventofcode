import spock.lang.Specification

class SquidGameSpec extends Specification {
    def "calculates winning board"() {
        expect:
        SquidGame.simulate('src/test/resources/squid-game.txt' as File) == 1924
    }
}
