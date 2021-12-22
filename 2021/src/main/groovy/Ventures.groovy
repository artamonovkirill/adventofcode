import commons.xy.Line
import commons.xy.Point

class Ventures {
    static Collection<Point> dangerous(File input) {
        def lines = input.readLines().collect { line(it) }
        println("Total lines: ${lines.size()}")
        lines.collectMany { it.points() }.groupBy { it }
                .findAll { it.value.size() > 1 }
                .keySet()
    }

    private static Line line(String input) {
        def coordinates = input.split('->')*.trim()*.split(',')
        new Line(
                start: new Point(x: coordinates[0][0] as int, y: coordinates[0][1] as int),
                end: new Point(x: coordinates[1][0] as int, y: coordinates[1][1] as int))
    }

    static void main(String... args) {
        println(dangerous('src/main/resources/ventures.txt' as File).size())
    }
}
