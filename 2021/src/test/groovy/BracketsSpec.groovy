import spock.lang.Specification

class BracketsSpec extends Specification {
    def 'solves example'() {
        given:
        def input = 'src/test/resources/brackets.txt' as File
        def solution = new Brackets(input)

        expect:
        solution.errors() == [')': 2, ']': 1, '>': 1, '}': 1]
        solution.errorsScore() == 26397
        solution.completions()*.join('') == ['}}]])})]', ')}>]})', '}}>}>))))', ']]}}]}]}>', '])}>']
        solution.completionScores() == [288957, 5566, 1480781, 995444, 294]
        solution.completionScore() == 288957
    }
}
