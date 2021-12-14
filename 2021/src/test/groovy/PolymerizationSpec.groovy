import spock.lang.Specification

class PolymerizationSpec extends Specification {
    def 'extends a template a few times'() {
        given:
        def rules = 'src/test/resources/polymerization.txt' as File
        def template = 'NNCB'
        def solution = new Polymerization(rules, template)

        expect:
        Polymerization.size(solution.process(1)) == 'NCNBCHB'.size()
        Polymerization.size(solution.process(2)) == 'NBCCNBBBCBHCB'.size()
        Polymerization.size(solution.process(3)) == 'NBBBCNCCNBBNBNBBCHBHHBCHB'.size()
        Polymerization.size(solution.process(4)) == 'NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB'.size()
        Polymerization.size(solution.process(5)) == 97
    }

    def 'counts template results after 10 steps'() {
        given:
        def rules = 'src/test/resources/polymerization.txt' as File
        def template = 'NNCB'
        def solution = new Polymerization(rules, template)

        when:
        def polymer = solution.process(10)

        then:
        Polymerization.size(polymer) == 3073
        solution.count(polymer) == [
                'B': 1749,
                'C': 298,
                'H': 161,
                'N': 865]
        solution.gap(polymer) == 1588
    }

    def 'counts template results after 40 steps'() {
        given:
        def rules = 'src/test/resources/polymerization.txt' as File
        def template = 'NNCB'
        def solution = new Polymerization(rules, template)

        when:
        def polymer = solution.process(40)

        then:
        solution.count(polymer)['B'] == 2192039569602
        solution.count(polymer)['H'] == 3849876073
        solution.gap(polymer) == 2188189693529
    }
}
