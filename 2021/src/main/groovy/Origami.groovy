import groovy.transform.MapConstructor
import groovy.transform.Memoized

class Fold {
    String axis
    int value
}

@MapConstructor
class Origami {
    List<List<Boolean>> page
    List<Fold> folds = []

    Origami(File input) {
        def lines = input.readLines()
        def coordinates = lines.takeWhile { it }*.split(',')*.collect { it as int }
        def width = coordinates.max { it[0] }[0]
        def height = coordinates.max { it[1] }[1]
        println("Page width $width, height $height")
        folds = lines.findAll {
            it.startsWith('fold along')
        }.collect {
            def fold = it.replace('fold along ', '').split('=')
            new Fold(axis: fold[0], value: fold[1] as int)
        }
        page = (0..height).collect { y ->
            (0..width).collect { x ->
                coordinates.contains([x, y])
            }
        }
    }

    String toString(String empty = ' ') {
        page.collect {
            it.collect { it ? '#' : empty }.join('')
        }.join('\n')
    }

    @Memoized
    Origami fold() {
        List<List<String>> page = folds.inject(page) { page, fold ->
            if (fold.axis == 'y') {
                println("Folding horizontally at ${fold.value}")
                foldHorizontally(page, fold)
            } else {
                println("Folding vertically at ${fold.value}")
                foldHorizontally(page.transpose(), fold).transpose()
            }
        }
        new Origami(page: page)
    }

    static def foldHorizontally(List<List<String>> page, Fold fold) {
        def upper = page.take(fold.value)
        def lower = page.drop(fold.value + 1)
        if (upper.size() < lower.size())
            throw new RuntimeException("Not implemented")

        def unchanged = upper.size() - lower.size()
        upper.take(unchanged) + [upper.drop(unchanged), lower.reverse()].transpose().collect {
            it.transpose().collect { a, b -> a || b }
        }
    }

    def points() {
        page.sum { it.count(true) }
    }

    static void main(String... args) {
        def input = 'src/main/resources/origami.txt' as File
        def origami = new Origami(input)
        println(origami.fold())
    }
}
