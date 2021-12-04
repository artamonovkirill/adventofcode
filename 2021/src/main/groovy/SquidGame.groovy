import groovy.transform.ToString

class SquidGame {
    static int simulate(File input) {
        def lines = input.readLines()
        def numbers = lines.head().split(',').collect { it as int }

        def boards = lines.tail()
                .findAll { it }
                .collate(5)
                .collect { new Board(it) }

        def loser = null
        def n = numbers.find { n ->
            boards = boards.findAll { b -> !b.check(n) }
            if (boards.size() == 1) loser = boards.head()
            !boards
        }
        n * loser.unmarked().sum { it.value }
    }

    static void main(String... args) {
        println(simulate('src/main/resources/squid-game.txt' as File))
    }
}

@ToString
class Field {
    int value
    boolean checked
}

@ToString
class Board {
    List<List<Field>> rows

    Board(List<String> rows) {
        this.rows = rows.collect { r ->
            r.trim().split(/[ ]+/).collect { f -> new Field(value: f as int) }
        }
    }

    boolean check(int number) {
        rows.each { r ->
            r.each { f ->
                if (f.value == number) {
                    f.checked = true
                }
            }
        }
        won()
    }

    boolean won() {
        rows.any { r -> r.every { f -> f.checked } } ||
                rows.transpose().any { c -> c.every { f -> f.checked } }
    }

    List<Field> unmarked() {
        rows.collectMany { r -> r.findAll { f -> !f.checked } }
    }
}
