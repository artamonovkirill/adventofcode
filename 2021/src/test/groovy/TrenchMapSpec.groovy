import spock.lang.Specification

class TrenchMapSpec extends Specification {

    def 'processes example image'() {
        given:
        def input = 'src/test/resources/trench/input.txt' as File
        def map = new TrenchMap(input)
        and:
        def source = 'src/test/resources/trench/0.txt' as File
        def first = 'src/test/resources/trench/1.txt' as File
        def second = 'src/test/resources/trench/2.txt' as File

        expect:
        TrenchMap.toString(TrenchMap.extend(map.process(0), 2)) == source.text
        TrenchMap.toString(map.process(1)) == first.text
        TrenchMap.toString(map.process(2)) ==
                TrenchMap.toString(TrenchMap.extend(second.readLines(), 2))
        TrenchMap.pixels(map.process(50)) == 3351
    }

}
