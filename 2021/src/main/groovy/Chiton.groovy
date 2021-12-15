import groovy.transform.EqualsAndHashCode

@EqualsAndHashCode
class CavePoint {
    int x
    int y

    int sum() {
        x + y
    }

    @Override
    String toString() {
        "[$x, $y]"
    }
}

@EqualsAndHashCode
class Path {
    int risk
    List<CavePoint> points

    @Override
    String toString() {
        "[$risk: $points]"
    }
}

class Chiton {
    List<List<Integer>> cave
    CavePoint last
    CavePoint first = new CavePoint(x: 0, y: 0)
    Map<CavePoint, Path> paths

    Chiton(File input) {
        this(input.readLines()*.collect { it as int })
    }

    Chiton(List<List<Integer>> cave) {
        this.cave = cave
        last = new CavePoint(x: cave[0].size() - 1, y: cave.size() - 1)
        paths = [(first): new Path(risk: 0, points: [first])]
    }

    @Override
    String toString() {
        cave*.join('').join('\n')
    }

    Integer risk(CavePoint point) {
        cave[point.y][point.x]
    }

    def best() {
        List<CavePoint> edge = [first]
        while (edge) {
            edge = edge.collectMany { p ->
                neighbours(p).findAll {
                    !paths[p].points.contains(it)
                }.collect { n ->
                    def current = paths[n]
                    def newRisk = risk(n) + paths[p].risk
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

    List<CavePoint> neighbours(CavePoint point) {
        [[0, -1], [-1, 0], [1, 0], [0, 1]].collect {
            x, y -> [x + point.x, y + point.y]
        }.findAll {
            x, y -> x >= 0 && y >= 0 && x <= last.x && y <= last.y
        }.collect {
            x, y -> new CavePoint(x: x, y: y)
        }
    }

    Chiton extend() {
        List<List<Integer>> extended = (0..4).collectMany { j ->
            cave.collect { row ->
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
