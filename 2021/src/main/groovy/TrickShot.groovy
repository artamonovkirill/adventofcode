import commons.xy.Point
import groovy.transform.Memoized

import static java.util.stream.Collectors.toList
import static java.util.stream.Stream.concat
import static java.util.stream.Stream.generate

class Target {
    int top
    int bottom
    int left
    int right
}

class TrickShot {
    @Memoized
    static List<List<Point>> all(Target target) {
        xs(target).collectMany {
            x -> ys(x, target)
        }
    }

    static int best(Target target) {
        all(target).collect {
            path -> path.collect { it.y }.max()
        }.max()
    }

    @Memoized
    static List<List<Integer>> xs(Target target) {
        def min = (1..target.left).find {
            (1..it).sum() >= target.left
        }
        (min..target.right).collect {
            accumulate(it..0)
        }.findAll {
            path -> path.any { x -> x >= target.left && x <= target.right }
        }
    }

    @Memoized
    static List<Integer> accumulate(List<Integer> xs) {
        xs.tail().inject([xs.head()]) { acc, e ->
            acc + [acc.last() + e]
        }
    }

    static List<List<Point>> ys(List<Integer> xs, Target target) {
        (target.bottom..target.right).collect { start ->
            def ys = accumulate(start..target.bottom)
            def extendedXs = concat(xs.stream(), generate(() -> xs.last()))
                    .limit(ys.size())
                    .collect(toList())
            [extendedXs, ys].transpose()
        }.findAll { path ->
            path.any { x, y -> x >= target.left && x <= target.right && y >= target.bottom && y <= target.top }
        }.collect {
            path -> path.collect { x, y -> new Point(x: x, y: y) }
        }
    }

    static void main(String... args) {
        def target = new Target(top: -72, bottom: -132, left: 155, right: 215)
        println(best(target))
        println(all(target).size())
    }

}
