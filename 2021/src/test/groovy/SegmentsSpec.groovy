import spock.lang.Specification

class SegmentsSpec extends Specification {
    def 'decodes'() {
        given:
        def input = 'src/test/resources/segments.txt' as File
        def readings = input.readLines()

        expect:
        Segments.decode(readings) == [8394, 9781, 1197, 9361, 4873, 8418, 4548, 1625, 8717, 4315]
    }
}
