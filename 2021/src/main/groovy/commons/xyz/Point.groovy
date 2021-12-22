package commons.xyz

import groovy.transform.EqualsAndHashCode

@EqualsAndHashCode
class Point {
    int x
    int y
    int z

    @Override
    String toString() {
        "[$x,$y,$z]"
    }
}
