import commons.xy.Point
import groovy.transform.EqualsAndHashCode
import groovy.transform.ToString

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

@ToString
@EqualsAndHashCode
class Line {
    Point start
    Point end

    def points() {
        if (start.x == end.x)
            (start.y..end.y).collect { new Point(x: start.x, y: it) }
        else if (start.y == end.y)
            (start.x..end.x).collect { new Point(x: it, y: start.y) }
        else {
            [(start.x..end.x), (start.y..end.y)].transpose().collect { x, y -> new Point(x: x, y: y) }
        }
    }
}
