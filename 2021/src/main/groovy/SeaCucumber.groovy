class SeaCucumber {
    static move(List<List<String>> field) {
        move(move(field, '>').transpose(), 'v').transpose()
    }

    static List<List<String>> move(List<List<String>> field, String type) {
        field.collect { row ->
            def moves = row.indexed().findAll { _, v ->
                v == type
            }.collectEntries { i, _ ->
                [(i): (i + 1) % row.size()]
            }.findAll { k, v ->
                row[v] == '.'
            }
            def from = moves.keySet()
            def to = moves.values()
            row.indexed().collect { i, v ->
                from.contains(i) ? '.' : to.contains(i) ? type : v
            }
        }
    }

    static List<List<String>> parse(String... input) {
        parse(input.toList())
    }

    static List<List<String>> parse(List<String> input) {
        input*.toList()
    }

    static int solve(List<List<String>> field, int step = 1) {
        def moved = move(field)
        if (moved == field) return step
        solve(moved, step + 1)
    }

    static void main(String... args) {
        def input = 'src/main/resources/cucumbers.txt' as File
        def field = parse(input.readLines())
        println(solve(field))
    }
}
