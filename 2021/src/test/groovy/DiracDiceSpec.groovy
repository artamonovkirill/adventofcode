import spock.lang.Ignore
import spock.lang.Specification

class DiracDiceSpec extends Specification {
    @Ignore
    def 'plays example game'() {
        expect:
        DiracDice.simulate(4, 8) == [444356092776315, 341960390180808]
    }
}
