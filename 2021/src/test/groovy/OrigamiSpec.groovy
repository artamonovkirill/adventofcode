import spock.lang.Specification

class OrigamiSpec extends Specification {
    def 'solves'() {
        given:
        def input = 'src/test/resources/origami/input.txt' as File
        def plain = 'src/test/resources/origami/plain.txt' as File
        def folded = 'src/test/resources/origami/folded.txt' as File

        when:
        def origami = new Origami(input)

        then:
        origami.toString('.') == plain.text
        origami.fold().toString('.') == folded.text
        origami.fold().points() == 16
    }
}
