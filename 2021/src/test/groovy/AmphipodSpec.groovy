import spock.lang.Specification

class AmphipodSpec extends Specification {
    def 'links are symmetrical'(List<String> link) {
        expect:
        Amphipod.LINKS[link[0]].contains(link[1])
        Amphipod.LINKS[link[1]].contains(link[0])

        where:
        link << Amphipod.LINKS.keySet().collectMany { k ->
            Amphipod.LINKS[k].collect { v -> [k, v] }
        }
    }

    def 'solves an example'() {
        given:
        def input = 'src/test/resources/amphipod/input.txt' as File

        expect:
        Amphipod.solve(input) == 44169
    }
}
