import commons.Point
import groovy.transform.Memoized

import java.util.stream.Stream

import static java.util.stream.Collectors.toList

class Target {
    int top
    int bottom
    int left
    int right
}

class TrickShot {
    static List<List<Point>> all(Target target) {
        xs(target.left, target.right).collectMany {
            x -> ys(x, target)
        }
    }

    @Memoized
    static int best(Target target) {
        all(target).collect {
            path -> path.collect { it.y }.max()
        }.max()
    }

    static List<List<Integer>> xs(int min, int max) {
        (1..max).collect { start ->
            Stream.iterate(start, x -> Math.max(x - 1, 0))
                    .limit(max * 5)
                    .collect(toList())
        }.collect {
            speeds -> accumulate(speeds)
        }.findAll {
            path -> path.any { x -> x >= min && x <= max }
        }
    }

    static List<Integer> accumulate(List<Integer> xs) {
        xs.tail().inject([xs.head()]) { acc, e ->
            acc + [acc.last() + e]
        }
    }

    static List<List<Point>> ys(List<Integer> xs, Target target) {
        (target.bottom..target.right).collect { start ->
            def speeds = Stream.iterate(start, e -> e - 1)
                    .takeWhile(e -> e >= target.bottom)
                    .collect(toList())
            def ys = accumulate(speeds)
            [xs, ys].transpose()
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
