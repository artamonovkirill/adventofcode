import spock.lang.Specification

class ALUSpec extends Specification {

    def 'checks sample input'() {
        when:
        ALU.check(13579246899999)

        then:
        noExceptionThrown()
    }

    def 'generates digits'() {
        expect:
        ALU.digits(99) == [9, 9]
    }

}
