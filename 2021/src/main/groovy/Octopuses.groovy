import commons.Matrix
import groovy.transform.EqualsAndHashCode
import groovy.transform.ToString

@ToString
@EqualsAndHashCode
class Octopus {
    int energy
    int x
    int y
}

class Octopuses extends Matrix {
    List<List<Octopus>> octopuses
    int flashes = 0
    int size

    Octopuses(File input) {
        super(input)
        octopuses = matrix.indexed().collect {
            j, row ->
                row.indexed().collect {
                    i, e -> new Octopus(energy: e as int, x: i, y: j)
                }
        }
        size = octopuses.sum { row -> row.size() }
    }


    static void main(String... args) {
        def input = 'src/main/resources/octopus.txt' as File
        println(new Octopuses(input).advance(100))
        println(new Octopuses(input).all())
    }

    int advance(int steps) {
        if (steps == 0) return this.flashes
        octopuses = octopuses.each { row ->
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
        List<Octopus> readyToFlash = octopuses.collectMany {
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
        neighbours(o.x, o.y, true).collect {
            p -> octopuses[p.y][p.x]
        }
    }

    int all(int steps = 1) {
        octopuses.each { row ->
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
