import groovy.transform.Memoized

import static java.lang.Integer.parseInt

class TrenchMap {
    String pattern
    List<String> image

    TrenchMap(File input) {
        def lines = input.readLines()
        pattern = lines.head()
        image = lines.drop(2)
    }

    static List<String> extend(List<String> image, int border, String value = '.') {
        def width = image.head().length() + border * 2
        def header = (1..border).collect {
            (1..width).collect { value }.join('')
        }
        header + image.collect {
            def prefix = (1..border).collect { value }.join('')
            prefix + it + prefix
        } + header
    }

    static String toString(List<String> image) {
        image.join('\n')
    }

    static int pixels(List<String> image) {
        image.sum { it.count('#') } as int
    }

    List<String> process(int steps) {
        process(extend(image, 3), steps)
    }

    @Memoized
    List<String> process(List<String> image, int steps) {
        if (steps == 0) return image
        if (steps % 10 == 0) println("$steps...")
        def width = image.head().size()
        def height = image.size()
        def result = (0..height - 3).collect { i ->
            def lines = image.drop(i).take(3)
            (0..width - 3).collect { j ->
                def square = lines.collect { line -> line.drop(j).take(3) }
                def binary = square.join('')
                        .replaceAll('#', '1')
                        .replaceAll('[.]', '0')
                def index = parseInt(binary, 2)
                pattern[index]
            }.join('')
        }
        process(extend(result, 3, result[0][0]), steps - 1)
    }

    static void main(String... args) {
        def input = 'src/main/resources/trench.txt' as File
        def map = new TrenchMap(input)
        def processed = map.process(50)
                .findAll { it.contains('#') }
        def left = processed*.indexOf('#').min() as int
        processed = processed.collect { it.drop(left) }
        def right = processed*.lastIndexOf('#').max() as int
        processed = processed.collect { it.take(right + 1) }
        println(toString(processed.findAll { it.contains('#') }))
        println(pixels(processed))
    }
}
