import spock.lang.Specification

class VenturesSpec extends Specification {
    def "calculates most dangerous places"() {
        expect:
        Ventures.dangerous('src/test/resources/ventures.txt' as File).size() == 12
    }
}
