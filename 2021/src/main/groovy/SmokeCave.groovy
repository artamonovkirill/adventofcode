import commons.Matrix
import groovy.transform.Memoized

class SmokeCave extends Matrix {
    SmokeCave(List<String> rows) {
        super(rows.collect { row -> row.collect { it as int } })
    }

    int risk() {
        lowPoints().sum { height, x, y -> height + 1 } as int
    }

    @Memoized
    List<List<Integer>> lowPoints() {
        matrix.indexed().collectMany { j, row ->
            row.indexed().findAll { i, e ->
                neighbours(i, j).every { height, x, y -> height > e }
            }.collect { i, e -> [e, i, j] }
        }
    }

    List<List<Integer>> neighbours(int i, int j) {
        super.neighbours(i, j).collect {
            [value(it), it.x, it.y]
        }
    }

    int basinsScore() {
        basins()*.size().sort { -it }.take(3).inject(1) { a, b -> a * b }
    }

    @Memoized
    List<List<List<Integer>>> basins() {
        lowPoints().collect {
            [it] + ups(it)
        }.collect { it.unique() }
    }

    List<List<Integer>> ups(List<Integer> point) {
        neighbours(point[1], point[2])
                .findAll { height, x, y -> height > point[0] && height < 9 }
                .collectMany { [it] + ups(it) }
    }

    static void main(String... args) {
        def input = 'src/main/resources/smoke-basin.txt' as File
        def cave = input.readLines()
        println(new SmokeCave(cave).risk())
        println(new SmokeCave(cave).basinsScore())
    }
}
