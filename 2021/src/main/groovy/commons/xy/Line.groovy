package commons.xy

import groovy.transform.EqualsAndHashCode
import groovy.transform.ToString

@ToString
@EqualsAndHashCode
class Line {
    Point start
    Point end

    def points() {
        if (start.x == end.x)
            (start.y..end.y).collect { new Point(x: start.x, y: it) }
        else if (start.y == end.y)
            (start.x..end.x).collect { new Point(x: it, y: start.y) }
        else {
            [(start.x..end.x), (start.y..end.y)].transpose().collect { x, y -> new Point(x: x, y: y) }
        }
    }
}
