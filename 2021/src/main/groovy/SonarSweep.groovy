class SonarSweep {
    static int increases(List<Integer> readings) {
        readings.collate(2, 1, false).count { a, b -> a < b }
    }

    static List<Integer> movingSum(List<Integer> readings, int window) {
        readings.collate(window, 1, false).collect { it.sum() }
    }

    static void main(String... args) {
        def report = new File('src/main/resources/sonar-sweep.txt')
        def readings = report.readLines().collect { it as int }
        println(increases(movingSum(readings, 3)))
    }
}
