import groovy.transform.EqualsAndHashCode

import static java.lang.Math.max
import static java.lang.Math.min

@EqualsAndHashCode
class Line {
    int start
    int end

    private Line(int start, int end) {
        this.start = min(start, end)
        this.end = max(start, end)
    }

    static of(int start, int end) {
        new Line(min(start, end), max(start, end))
    }

    static Optional<Line> intersect(Line a, Line b) {
        b.end < a.start || b.start > a.end
                ? Optional.empty()
                : Optional.of(new Line(max(a.start, b.start), min(a.end, b.end)))
    }

    List<Line> cut(Line other) {
        intersect(this, other)
                .map({ intersection ->
                    def remainder = []
                    if (intersection.start > start) remainder << new Line(start, intersection.start - 1)
                    if (intersection.end < end) remainder << new Line(intersection.end + 1, end)
                    return remainder
                }).orElse([this])

    }

    @Override
    String toString() {
        "$start -> $end"
    }

    long length() {
        end - start + 1
    }
}

@EqualsAndHashCode
class Cuboid {
    Line x
    Line y
    Line z

    long size() {
        x.length() * y.length() * z.length()
    }

    @Override
    String toString() {
        "[x: $x, y: $y, z: $z]"
    }

    static Optional<Cuboid> intersect(Cuboid a, Cuboid b) {
        Line.intersect(a.x, b.x).flatMap {
            x ->
                Line.intersect(a.y, b.y).flatMap {
                    y ->
                        Line.intersect(a.z, b.z).map {
                            z -> new Cuboid(x: x, y: y, z: z)
                        }
                }
        }
    }

    List<Cuboid> cut(Cuboid other) {
        def x = this.x.cut(other.x)
        def y = this.y.cut(other.y)
        def z = this.z.cut(other.z)
        if (x.size() > 0) {
            return x.collect {
                new Cuboid(x: it, y: this.y, z: this.z)
            } + new Cuboid(x: other.x, y: this.y, z: this.z).cut(other)
        }
        if (y.size() > 0) {
            return y.collect {
                new Cuboid(x: this.x, y: it, z: this.z)
            } + new Cuboid(x: this.x, y: other.y, z: this.z).cut(other)
        }
        if (z.size() > 0) {
            return z.collect {
                new Cuboid(x: this.x, y: this.y, z: it)
            } + new Cuboid(x: this.x, y: this.y, z: other.z).cut(other)
        }
        return []
    }
}

class ReactorReboot {
    static final BOUNDARY = new Cuboid(
            x: Line.of(-50, 50),
            y: Line.of(-50, 50),
            z: Line.of(-50, 50))

    static long process(File input, boolean init = true) {
        input.readLines().inject([]) { cuboids, command ->
            def (action, lines) = command.split(' ')
            parse(lines, init).map { cuboid ->
                println("turning $action $cuboid")
                def result = action == 'on' ? [cuboid] : []
                result + cut(cuboids, cuboid)
            }.orElse(cuboids)
        }.sum { it.size() }
    }

    static List<Cuboid> cut(List<Cuboid> cuboids, Cuboid cuboid) {
        cuboids.collectMany { c ->
            Cuboid.intersect(c, cuboid).map { intersection ->
                c.cut(intersection)
            }.orElse([c])
        }
    }

    static Optional<Cuboid> parse(String line, boolean init = true) {
        def (x, y, z) = ['x', 'y', 'z'].collect {
            line.find("$it=[-]?[0-9]+[.]{2}[-]?[0-9]+")
                    .replace("$it=", '')
                    .split('[.]{2}')
                    *.toInteger()
        }.collect {
            start, end -> Line.of(start, end)
        }
        def cuboid = new Cuboid(x: x, y: y, z: z)
        init ? Cuboid.intersect(cuboid, BOUNDARY) : Optional.of(cuboid)
    }

    static void main(String... args) {
        def input = 'src/main/resources/reactor.txt' as File
        println(process(input, false))
    }
}
