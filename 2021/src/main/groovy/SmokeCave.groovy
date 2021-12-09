import groovy.transform.Memoized

class SmokeCave {
    final List<List<Integer>> matrix

    SmokeCave(List<String> rows) {
        matrix = rows.collect { row -> row.collect { it as int } }
    }

    int risk() {
        lowPoints().sum { height, x, y -> height + 1 }
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
        [[0, -1], [-1, 0], [1, 0], [0, 1]].collect {
            x, y -> [x + i, y + j]
        }.findAll {
            x, y -> i >= 0 && y >= 0 && x < matrix[j].size() && y < matrix.size()
        }.collect {
            x, y -> [matrix[y][x], x, y]
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
