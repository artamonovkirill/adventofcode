import static java.lang.Integer.parseInt
import static java.lang.Long.parseLong

class Diagnostic {
    static String gamma(List<String> report) {
        report.collect { it.findAll() }.transpose()
                .collect { bits -> bits.count { b -> b == '1' } > bits.size() / 2 ? 1 : 0 }
                .join('')
    }

    static String epsilon(String gamma) {
        gamma.collect { it == '1' ? '0' : '1' }.join('')
    }

    static int power(String gamma, String epsilon) {
        parseInt(gamma, 2) * parseInt(epsilon, 2)
    }

    static long oxygen(List<String> report) {
        parseLong(value(report,
                { zeros, ones -> ones == zeros ? '1' : ones > zeros ? '1' : '0' },
                0), 2)
    }

    static long co2(List<String> report) {
        parseLong(value(report,
                { zeros, ones -> ones == zeros ? '0' : ones < zeros ? '1' : '0' },
                0), 2)
    }

    private static String value(List<String> report, Closure<String> comparator, int n) {
        if (report.size() == 1)
            return report.head()

        def bits = report
                .collect { it[n] }
                .groupBy { it }
                .collectEntries { [(it.key): it.value.size()] }
        def bit = comparator(bits['0'], bits['1'])
        value(report.findAll { it[n] == bit }, comparator, n + 1)
    }

    static void main(String... args) {
        def report = new File('src/main/resources/diagnostic.txt')
        def readings = report.readLines()
        def gamma = gamma(readings)
        def epsilon = epsilon(gamma)
        println("Power: ${power(gamma, epsilon)}")
        def oxygen = oxygen(readings)
        def co2 = co2(readings)
        println("Life: ${oxygen * co2}")
    }
}
