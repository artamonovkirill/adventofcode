import spock.lang.Specification

class DiagnosticSpec extends Specification {
    def 'calculates rates'() {
        given:
        def input = [
                '00100', '11110', '10110', '10111', '10101', '01111',
                '00111', '11100', '10000', '11001', '00010', '01010']

        when:
        def gamma = Diagnostic.gamma(input)
        def epsilon = Diagnostic.epsilon(gamma)

        then:
        gamma == '10110'
        epsilon == '01001'
        Diagnostic.power(gamma, epsilon) == 198
        and:
        Diagnostic.oxygen(input) == 23
        Diagnostic.co2(input) == 10
    }
}
