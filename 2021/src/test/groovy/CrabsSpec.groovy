import spock.lang.Specification

class CrabsSpec extends Specification {
    def 'calculates best position'() {
        expect:
        Crabs.best([16,1,2,0,4,2,7,1,2,14]) == [5, 168]
    }
}
