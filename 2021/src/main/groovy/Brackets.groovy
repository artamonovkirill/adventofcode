import groovy.transform.Memoized

class Brackets {
    static pairs = ['(': ')', '{': '}', '[': ']', '<': '>']
    static errorScores = [')': 3, ']': 57, '}': 1197, '>': 25137]
    static completionScores = [')': 1, ']': 2, '}': 3, '>': 4]
    List<List<String>> lines

    Brackets(File input) {
        lines = input.readLines()*.toList()
    }

    @Memoized
    Map<String, Integer> errors() {
        lines.collect { parse(it) }
                .collect { _, error -> error }
                .findAll { it }
                .groupBy { it }
                .collectEntries { e -> [e.key, e.value.size()] }
    }

    @Memoized
    def completions() {
        lines.collect { parse(it) }
                .findAll { stack, _ -> stack }
                .collect { stack, _ -> stack }
                .collect { stack -> stack.collect { c -> pairs[c] } }
    }

    @Memoized
    static parse(List<String> characters) {
        def stack = [characters.head()]
        for (String c : characters.tail()) {
            if (!stack || pairs.keySet().contains(c))
                stack = [c] + stack
            else {
                if (pairs[stack.head()] == c)
                    stack = stack.tail()
                else
                    return [null, c]
            }
        }
        return [stack, null]
    }

    int errorsScore() {
        errors().collect { e -> errorScores[e.key] * e.value }.sum()
    }

    static void main(String... args) {
        def brackets = new Brackets('src/main/resources/brackets.txt' as File)
        println(brackets.errorsScore())
        println(brackets.completionScore())
    }

    @Memoized
    List<Long> completionScores() {
        completions()*.inject(0L) { acc, e -> acc * 5 + completionScores[e] }
    }

    long completionScore() {
        completionScores().sort()[completionScores().size() / 2]
    }
}
