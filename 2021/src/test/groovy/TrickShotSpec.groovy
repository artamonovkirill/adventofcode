import spock.lang.Specification

class TrickShotSpec extends Specification {
    def 'solves example'() {
        given:
        def target = new Target(top: -5, bottom: -10, left: 20, right: 30)

        expect:
        TrickShot.best(target) == 45
        TrickShot.all(target).size() == 112
    }
}
