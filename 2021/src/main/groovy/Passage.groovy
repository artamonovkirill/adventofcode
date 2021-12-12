import groovy.transform.Memoized

class Passage {
    Map<String, List<String>> caves

    Passage(List<String> input) {
        def connections = input.collect { it.split('-') }*.toList()
        caves = connections.collectMany {
            it
        }.unique().collectEntries { c ->
            def from = connections.findAll {
                a, b -> a == c
            }*.get(1)
            def to = connections.findAll {
                a, b -> b == c
            }*.get(0)
            [(c): from + to]
        }.findAll {
            e -> e.key != 'end'
        }.collectEntries {
            e -> [(e.key): e.value.findAll { it != 'start' }]
        }
    }

    Set<String> paths() {
        def start = [['start']].toSet()
        extendAll(start).findAll { it.last() == 'end' }*.join(',')
    }

    Set<List<String>> extendAll(Set<List<String>> paths) {
        def groups = paths
                .collect { extendOne(it) }
                .findAll { it }
                .collectMany { it }
                .groupBy { it.last() == 'end' }

        def finalists = groups.get(true, []).toSet()
        def candidates = groups.get(false, []).toSet()
        if (!candidates) return finalists

        println("Found ${finalists.size()} paths, ${candidates.size()} more candidates")
        finalists + extendAll(candidates)
    }

    @Memoized
    Set<List<String>> extendOne(List<String> path) {
        caves[path.last()].collect {
            if (it.matches('[a-z]{1,2}')) {
                def lowers = path.findAll {
                    it.matches('[a-z]{1,2}')
                }.countBy { it }
                if (lowers.values().max() == 2 && lowers[it])
                    return null
            }
            path + [it]
        }.findAll { it }.toSet()
    }

    static void main(String... args) {
        def input = 'src/main/resources/passage.txt' as File
        def connections = input.readLines()
        println(new Passage(connections).paths().size())
    }
}
