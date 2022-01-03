import spock.lang.Specification

class SeaCucumberSpec extends Specification {

    def 'advances simple eastward example'() {
        given:
        def field = SeaCucumber.parse(input)

        expect:
        SeaCucumber.move(field) == SeaCucumber.parse(output)

        where:
        input         | output
        '...>>>>>...' | '...>>>>.>..'
        '...>>>>.>..' | '...>>>.>.>.'
    }

    def 'advances simple southward example'() {
        given:
        def field = SeaCucumber.parse(
                '..........',
                '.>v....v..',
                '.......>..',
                '..........')

        expect:
        SeaCucumber.move(field) == SeaCucumber.parse(
                '..........',
                '.>........',
                '..v....v>.',
                '..........')
    }

    def 'advances medium example'() {
        given:
        def field = SeaCucumber.parse(
                '...>...',
                '.......',
                '......>',
                'v.....>',
                '......>',
                '.......',
                '..vvv..')

        expect:
        (1..4).inject(field) { f, _ ->
            SeaCucumber.move(f)
        } == SeaCucumber.parse(
                '>......',
                '..v....',
                '..>.v..',
                '.>.v...',
                '...>...',
                '.......',
                'v......')
    }

    def 'finds a safe spot'() {
        given:
        def field = SeaCucumber.parse(
                'v...>>.vv>',
                '.vv>>.vv..',
                '>>.>v>...v',
                '>>v>>.>.v.',
                'v>v.vv.v..',
                '>.>>..v...',
                '.vv..>.>v.',
                'v.v..>>v.v',
                '....v..v.>')

        expect:
        SeaCucumber.solve(field) == 58
    }

}
