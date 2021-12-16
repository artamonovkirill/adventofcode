import groovy.transform.EqualsAndHashCode
import commons.Matrix
import commons.Point

@EqualsAndHashCode
class Path {
    int risk
    List<Point> points

    @Override
    String toString() {
        "[$risk: $points]"
    }
}

class Chiton extends Matrix {
    Point first = new Point(x: 0, y: 0)
    Map<Point, Path> paths

    Chiton(File input) {
        this(input.readLines()*.collect { it as int })
    }

    Chiton(List<List<Integer>> cave) {
        super(cave)
        paths = [(first): new Path(risk: 0, points: [first])]
    }

    def best() {
        List<Point> edge = [first]
        while (edge) {
            edge = edge.collectMany { p ->
                neighbours(p).findAll {
                    !paths[p].points.contains(it)
                }.collect { n ->
                    def current = paths[n]
                    def newRisk = value(n) + paths[p].risk
                    if (!current || newRisk < current.risk) {
                        paths[n] = new Path(risk: newRisk, points: paths[p].points + [n])
                        return n
                    }
                    return null
                }
            }.findAll { it }
            println("Edge is ${edge.size()} big")
        }
        return paths[last].risk
    }

    Chiton extend() {
        List<List<Integer>> extended = (0..4).collectMany { j ->
            matrix.collect { row ->
                (0..4).collectMany { i ->
                    row.collect { e ->
                        e + i + j
                    }.collect {
                        e -> e > 9 ? e - 9 : e
                    }
                }
            }
        }
        new Chiton(extended)
    }

    static void main(String... args) {
        def input = new File('src/main/resources/chiton.txt')
        def solution = new Chiton(input)
        def extended = solution.extend()
        println(extended.best())
    }
}
