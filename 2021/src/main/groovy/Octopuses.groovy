import groovy.transform.EqualsAndHashCode
import groovy.transform.ToString

@ToString
@EqualsAndHashCode
class Octopus {
    int energy
    int x
    int y
}

class Octopuses {
    List<List<Octopus>> matrix
    int flashes = 0
    int size

    Octopuses(File input) {
        matrix = input.readLines().indexed().collect {
            j, line ->
                line.toList().indexed().collect {
                    i, e -> new Octopus(energy: e as int, x: i, y: j)
                }
        }
        size = matrix.sum { row -> row.size() }
    }


    static void main(String... args) {
        def input = 'src/main/resources/octopus.txt' as File
        println(new Octopuses(input).advance(100))
        println(new Octopuses(input).all())
    }

    int advance(int steps) {
        if (steps == 0) return this.flashes
        matrix = matrix.each { row ->
            row.each {
                o -> o.energy++
            }
        }

        def flashes = flash([])
        flashes.each { o -> o.energy = 0 }
        this.flashes += flashes.size()

        advance(steps - 1)
    }

    List<Octopus> flash(List<Octopus> alreadyFlashed) {
        List<Octopus> readyToFlash = matrix.collectMany {
            row -> row.findAll { it.energy > 9 }
        }.findAll {
            !alreadyFlashed.contains(it)
        }
        if (!readyToFlash) return alreadyFlashed
        readyToFlash.each {
            o ->
                neighbours(o).findAll {
                    !alreadyFlashed.contains(it)
                }.each {
                    n -> n.energy++
                }
        }
        return flash(alreadyFlashed + readyToFlash)
    }

    List<Octopus> neighbours(Octopus o) {
        [[0, -1], [-1, -1], [-1, 0], [-1, 1], [1, 0], [1, 1], [0, 1], [1, -1]].collect {
            x, y -> [x + o.x, y + o.y]
        }.findAll {
            x, y -> x >= 0 && y >= 0 && x < matrix[o.y].size() && y < matrix.size()
        }.collect {
            x, y -> matrix[y][x]
        }
    }

    int all(int steps = 1) {
        matrix.each { row ->
            row.each {
                o -> o.energy++
            }
        }

        def flashes = flash([])
        flashes.each { o -> o.energy = 0 }

        if (flashes.size() == size)
            return steps

        all(steps + 1)
    }
}
