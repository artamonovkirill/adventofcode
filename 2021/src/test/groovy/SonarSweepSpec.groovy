import spock.lang.Specification

class SonarSweepSpec extends Specification {
    def 'calculates rises'() {
        given:
        def report = [
                199,
                200,
                208,
                210,
                200,
                207,
                240,
                269,
                260,
                263]

        expect:
        SonarSweep.increases(report) == 7
    }

    def 'calculates sliding sum'() {
        given:
        def report = [
                199,
                200,
                208,
                210,
                200,
                207,
                240,
                269,
                260,
                263]

        expect:
        SonarSweep.movingSum(report, 3) == [
                607,
                618,
                618,
                617,
                647,
                716,
                769,
                792]
    }
}
