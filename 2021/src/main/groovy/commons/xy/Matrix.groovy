package commons.xy

class Matrix {
    List<List<Integer>> matrix
    Point last

    Matrix(File input) {
        this(input.readLines()*.collect { it as int })
    }

    Matrix(List<List<Integer>> matrix) {
        this.matrix = matrix
        last = new Point(x: matrix[0].size() - 1, y: matrix.size() - 1)
    }

    List<Point> neighbours(Point point, boolean diagonal = false) {
        neighbours(point.x, point.y, diagonal)
    }

    List<Point> neighbours(int i, int j, boolean diagonal = false) {
        def neighbours = [[0, -1], [-1, 0], [1, 0], [0, 1]]
        if (diagonal) neighbours += [[-1, -1], [-1, 1], [1, 1], [1, -1]]
        neighbours.collect {
            x, y -> [x + i, y + j]
        }.findAll {
            x, y -> x >= 0 && y >= 0 && x <= last.x && y <= last.y
        }.collect {
            x, y -> new Point(x: x, y: y)
        }
    }

    int value(Point p) {
        matrix[p.y][p.x]
    }

    @Override
    String toString() {
        matrix*.join('').join('\n')
    }
}
