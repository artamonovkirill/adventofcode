import spock.lang.Specification

class TrickShotSpec extends Specification {
    def 'finds best direction'() {
        given:
        def target = new Target(top: -5, bottom: -10, left: 20, right: 30)

        expect:
        TrickShot.best(target) == 45
    }

    def 'finds all shots'() {
        given:
        def target = new Target(top: -5, bottom: -10, left: 20, right: 30)

        when:
        def shots = TrickShot.all(target)

        then:
        shots.size() == 112
    }
}
