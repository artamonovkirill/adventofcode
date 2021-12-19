import commons.d3.Point

class BeaconScanner {
    static List<Point> rotateByX(Point p) {
        // [X, Y, Z]
        // [X, Z, -Y]
        // [X, -Y, -Z]
        // [X, -Z, Y]
        [p,
         new Point(x: p.x, y: p.z, z: -p.y),
         new Point(x: p.x, y: -p.y, z: -p.z),
         new Point(x: p.x, y: -p.z, z: p.y)]
    }

    static List<Point> rotateByY(Point p) {
        // [X, Y, Z]
        // [Z, Y, -X]
        // [-X, Y, -Z]
        // [-Z, Y, X]
        [
                p,
                new Point(x: p.z, y: p.y, z: -p.x),
                new Point(x: -p.x, y: p.y, z: -p.z),
                new Point(x: -p.z, y: p.y, z: p.x)
        ]
    }

    static List<Point> rotateByZ(Point p) {
        // [X, Y, Z]
        // [Y, -X, Z]
        // [-X, -Y, Z]
        // [-Y, X, Z]
        [
                p,
                new Point(x: p.y, y: -p.x, z: p.z),
                new Point(x: -p.x, y: -p.y, z: p.z),
                new Point(x: -p.y, y: p.x, z: p.z)
        ]
    }

    static List<Point> rotate(Point p) {
        rotateByX(p).collectMany { x ->
            rotateByY(x).collectMany { y ->
                rotateByZ(y)
            }
        }
    }

    static Set<List<Point>> rotate(List<Point> ps) {
        ps.collect { rotate(it) }.transpose().toSet()
    }

    static List<List<Point>> read(File input) {
        input.text.split('--- scanner [0-9]+ ---\n').findAll {
            it
        }.collect { reading ->
            reading.split('\n').findAll {
                it
            }.collect { point ->
                def (x, y, z) = point.split(',').collect { it as int }
                new Point(x: x, y: y, z: z)
            }
        }
    }

    static intersect(List<List<Point>> scanners) {
        def beacons = scanners.head().toSet()
        def newlyRotatedScanners = [0: scanners.head()]
        def scannerPositions = []
        Map<Integer, List<Point>> unmatchedScanners = scanners.indexed().drop(1).collectEntries {
            i, s -> [(i): s]
        }
        (0..unmatchedScanners.size()).each {
            if (unmatchedScanners) {
                def matched = [:]
                for (int i : unmatchedScanners.keySet()) {
                    for (Map.Entry<Integer, List<Point>> r : newlyRotatedScanners) {
                        def match = match(unmatchedScanners[i], r.value)
                        if (match) {
                            def (scanner, points) = match
                            println("Found a match with scanner #${r.key} for scanner #${i}")
                            scannerPositions.add(scanner)
                            beacons.addAll(points)
                            matched[i] = points
                            break
                        }
                    }
                }
                matched.keySet().each { unmatchedScanners.remove(it) }
                newlyRotatedScanners = matched
                println("Unmatched scanners: ${unmatchedScanners.keySet()}")
                println("Rotated scanners: ${newlyRotatedScanners.keySet()}")
            }
        }
        [beacons, scannerPositions]
    }

    static match(List<Point> scanner, List<Point> base) {
        def rotations = rotate(scanner)
        assert rotations.size() == 24
        for (List<Point> rotation : rotations) {
            for (Point p : rotation) {
                for (Point b : base) {
                    def zero = new Point(
                            x: b.x - p.x,
                            y: b.y - p.y,
                            z: b.z - p.z)
                    def rotated = rotation.collect {
                        new Point(
                                x: it.x + zero.x,
                                y: it.y + zero.y,
                                z: it.z + zero.z)
                    }

                    def matches = rotated.findAll { base.contains(it) }
                    if (matches.size() >= 12) {
                        return [zero, rotated]
                    }
                }
            }
        }
        null
    }

    static void main(String... args) {
        def input = 'src/main/resources/scanners.txt' as File
        def scanners = read(input)
        def (beacons, scannerPositions) = intersect(scanners)
        println(beacons.size())
        println(furthest(scannerPositions))
    }

    static furthest(List<Point> scanners) {
        def max = 0
        scanners.each { a ->
            scanners.each { b ->
                def distance = Math.abs(a.x - b.x) + Math.abs(a.y - b.y) + Math.abs(a.z - b.z)
                if (distance > max) max = distance
            }
        }
        max
    }

}
