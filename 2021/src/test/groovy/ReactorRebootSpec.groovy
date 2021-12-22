import spock.lang.Specification

class ReactorRebootSpec extends Specification {
    def 'solves a simple example'() {
        given:
        def input = 'src/test/resources/reactor/simple.txt' as File

        expect:
        ReactorReboot.process(input) == 27 + 19 - 8 + 1
    }

    def 'solves a larger example'() {
        given:
        def input = 'src/test/resources/reactor/larger.txt' as File

        expect:
        ReactorReboot.process(input) == 590784
    }

    def 'solves a complex example'() {
        given:
        def input = 'src/test/resources/reactor/complex.txt' as File

        expect:
        ReactorReboot.process(input, false) == 2758514936282235
    }
}
