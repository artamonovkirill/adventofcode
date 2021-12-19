package commons.d3

import groovy.transform.EqualsAndHashCode

@EqualsAndHashCode
class Point {
    int x
    int y
    int z

    @Override
    String toString() {
        "[x: $x, y: $y, z: $z]"
    }
}
