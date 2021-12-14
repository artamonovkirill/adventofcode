import groovy.transform.Memoized

class Polymerization {
    Map<String, List<String>> rules
    String first
    String last
    String template

    Polymerization(File rules, String template) {
        this.rules = rules.readLines()*.split(' -> ').collectEntries {
            def pair = it[0]
            [(pair): [pair[0] + it[1], it[1] + pair[1]]]
        }
        this.template = template
        first = template[0]
        last = template.toList().last()
    }

    Map<String, Long> process(int steps) {
        Map<String, Long> polymer = template
                .toList()
                .collate(2, 1, false)
                .collect {
                    it.join('')
                }.countBy { it }
        process(polymer, steps)
    }

    Map<String, Long> process(Map<String, Long> template, int steps) {
        if (steps == 0) return template
        Map<String, Long> polymer = [:]
        template.each { pair ->
            rules[pair.key].each {
                polymer[it] = polymer.getOrDefault(it, 0 as long) + pair.value
            }
        }
        return process(polymer, steps - 1)
    }

    static long size(Map<String, Long> polymer) {
        polymer.values().sum() + 1
    }

    @Memoized
    Map<String, Long> count(Map<String, Long> polymer) {
        Map<String, Long> counts = polymer.keySet().collectMany {
            it.toList()
        }.unique().collectEntries {
            [(it): 0 as long]
        }
        polymer.each { k, v ->
            k.each {
                l -> counts[l] += v
            }
        }
        counts[first] += 1
        counts[last] += 1
        return counts.collectEntries {
            k, v -> [(k): v / 2]
        }
    }

    long gap(Map<String, Long> polymer) {
        def counts = count(polymer)
        counts.max { it.value }.value - counts.min { it.value }.value
    }

    static void main(String... args) {
        def rules = 'src/main/resources/polymerization.txt' as File
        def template = 'OOBFPNOPBHKCCVHOBCSO'
        def solution = new Polymerization(rules, template)
        def polymer = solution.process(40)
        println(solution.gap(polymer))
    }
}
