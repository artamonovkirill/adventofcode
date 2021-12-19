import commons.d3.Point
import spock.lang.Specification

class BeaconScannerSpec extends Specification {
    def 'generates orientations'() {
        expect:
        BeaconScanner.rotate(new Point(x: 0, y: 0, z: 0)).size() == 64
        BeaconScanner.rotate(new Point(x: 1, y: 2, z: 3)).size() == 64
        BeaconScanner.rotate(new Point(x: 1, y: 2, z: 3)).toSet().size() == 24
    }

    def 'parses scanner readings'() {
        given:
        def input = 'src/test/resources/beacon/rotations.txt' as File

        when:
        def scanners = BeaconScanner.read(input)

        then:
        scanners.size() == 5
        scanners.every { it.size() == 6 }
    }

    def 'rotates a scanner'() {
        given:
        def input = 'src/test/resources/beacon/rotations.txt' as File
        def scanners = BeaconScanner.read(input)

        when:
        def rotations = BeaconScanner.rotate(scanners[0])

        then:
        scanners.indexed().each { i, s ->
            println("Comparing #$i")
            assert rotations.contains(s)
        }
    }

    def 'solves an example'() {
        given:
        def input = 'src/test/resources/beacon/example.txt' as File
        def scanners = BeaconScanner.read(input)

        when:
        def (beacons, scannerPositions) = BeaconScanner.intersect(scanners)

        then:
        beacons.size() == 79
        BeaconScanner.furthest(scannerPositions) == 3621
    }
}