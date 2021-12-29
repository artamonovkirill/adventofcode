import groovy.transform.Memoized

class Amphipod {
    static final ENTRANCES = ['A': 'H2', 'B': 'H4', 'C': 'H6', 'D': 'H8']
    static final PRICE = ['A': 1, 'B': 10, 'C': 100, 'D': 1000]
    static final LINKS = links()

    static Map<String, List<String>> links() {
        Map<String, List<String>> result = [:]
        (0..10).each { result["H$it".toString()] = [] }
        (0..9).each {
            result["H$it"] << "H${it + 1}".toString()
        }
        (1..10).each {
            result["H$it"] << "H${it - 1}".toString()
        }

        ['A', 'B', 'C', 'D'].each {
            result["${it}0".toString()] = ["${it}1".toString(), ENTRANCES[it]]
            result[ENTRANCES[it]] << "${it}0".toString()
            result["${it}1".toString()] = ["${it}0".toString(), "${it}2".toString()]
            result["${it}2".toString()] = ["${it}1".toString(), "${it}3".toString()]
            result["${it}3".toString()] = ["${it}2".toString()]
        }
        return result
    }

    @Memoized
    static List<List<String>> paths(String name, String position) {
        def paths = extend([position])*.tail().findAll {
            it && !ENTRANCES.values().contains(it.last())
        }.findAll {
            it.last().startsWith('H') || it.last().startsWith(name)
        }
        position.startsWith('H') ? paths.findAll {
            it.last().startsWith(name)
        } : paths
    }

    static paths(Map<String, String> board) {
        board.findAll { position, name ->
            "${name}3" != position
        }.findAll { position, name ->
            position != "${name}2" || board["${name}3"] != name
        }.findAll { position, name ->
            position != "${name}1" || board["${name}2"] != name
        }.findAll { position, name ->
            position != "${name}0" || board["${name}1"] != name
        }.collectMany { position, name ->
            paths(name, position, board).collect { path ->
                [name, position, path]
            }
        }
    }

    @Memoized
    static List<List<String>> paths(String name, String position, Map<String, String> board) {
        def occupied = board.keySet()
        paths(name, position).findAll { p ->
            p.every { !occupied.contains(it) }
        }.findAll { p ->
            def last = p.last()
            if (last.startsWith('H')) return true
            if (last == "${name}3") return true
            if (last == "${name}2" && board["${name}3"] == name) return true
            if (last == "${name}1" && board["${name}2"] == name && board["${name}3"] == name) return true
            if (last == "${name}0" && board["${name}1"] == name && board["${name}2"] == name && board["${name}3"] == name) return true
            false
        }
    }

    @Memoized
    static List<List<String>> extend(List<String> path) {
        [path] + LINKS[path.last()].findAll {
            !path.contains(it)
        }.collectMany {
            extend(path + [it])
        }
    }

    static solve(File input) {
        def data = input.readLines()
        def start = parse(data)
        def finals = possibleMoves([(start): [[], 0]], [:])
        def (steps, cost) = finals.min { steps, cost -> cost }.value
        println(steps)
        cost
    }

    @Memoized
    static Map<String, String> parse(List<String> move) {
        def (hall, level1, level2, level3, level4) = move.subList(1, 6)
        Map<String, String> amphipods = hall.substring(1, 12).toList().indexed().findAll {
            _, v -> v =~ /[ABCD]/
        }.collectEntries { i, v ->
            [("H$i".toString()): v]
        }

        ['A': 3, 'B': 5, 'C': 7, 'D': 9].each { l, i ->
            if (level1[i] =~ /[ABCD]/) amphipods["${l}0".toString()] = level1[i]
            if (level2[i] =~ /[ABCD]/) amphipods["${l}1".toString()] = level2[i]
            if (level3[i] =~ /[ABCD]/) amphipods["${l}2".toString()] = level3[i]
            if (level4[i] =~ /[ABCD]/) amphipods["${l}3".toString()] = level4[i]
        }

        assert amphipods.size() == 16
        return amphipods.sort()
    }

    static Map<Map<String, String>, List> possibleMoves(
            Map<Map<String, String>, List> boards,
            Map<Map<String, String>, List> finals,
            int step = 0) {
        println("#$step: processing ${boards.size()} boards, already found ${finals.size()} solutions")
        Map<Map<String, String>, List> result = [:]
        int discarded = 0

        boards.each { b, acc ->
            def (steps, cost) = acc
            def occupied = b.keySet()
            assert occupied.size() == 16

            def paths = paths(b)
            if (!paths) {
                if (occupied.any { it.startsWith('H') }) {
                    discarded++
                } else {
                    if (!finals[b] || finals[b][1] > cost) {
                        finals[b] = [steps, cost]
                    }
                }
            }

            paths.each { name, position, path ->
                def price = PRICE[name]
                def moveCost = cost + price * path.size()
                def m = move(b, name, position, path.last())
                if (!result[m] || result[m][1] > moveCost)
                    result[m] = [steps + [[name, position, path.last()]], moveCost]
            }
        }

        println("Discarded $discarded solutions")

        if (!result)
            return finals

        possibleMoves(result, finals, step + 1)
    }

    static Map<String, String> move(Map<String, String> board, String letter, String from, String to) {
        assert letter == board[from]
        board.collectEntries { position, name ->
            [(position == from ? to : position): name]
        }.sort()
    }

    static void main(String... args) {
        def input = 'src/main/resources/amphipod/input.txt' as File
        println(solve(input))
    }

}