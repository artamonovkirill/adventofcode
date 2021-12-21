package commons.xy

import groovy.transform.EqualsAndHashCode

@EqualsAndHashCode
class Point {
    int x
    int y

    @Override
    String toString() {
        "[$x, $y]"
    }
}
