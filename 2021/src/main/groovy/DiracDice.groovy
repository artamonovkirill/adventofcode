import groovy.transform.EqualsAndHashCode
import groovy.transform.Memoized
import groovy.transform.ToString

import static java.lang.Math.max

@ToString
@EqualsAndHashCode
class Universe {
    int firstPosition
    int secondPosition
    int firstScore
    int secondScore
}

class DiracDice {

    static simulate(int firstPosition, int secondPosition) {
        simulate([new Universe(
                firstPosition: firstPosition,
                secondPosition: secondPosition,
                firstScore: 0,
                secondScore: 0): 1L], 0L, 0L)
    }

    static simulate(Map<Universe, Long> universes, long firstWins, long secondWins) {
        println("Universes before first roll: ${universes.values().sum()}")

        if (!universes)
            return [firstWins, secondWins]

        Map<Universe, Long> newUniverses = [:]
        universes.each { universe, count ->
            splits().each { dice ->
                def (int score, int position) = roll(universe.firstScore, universe.firstPosition, dice)
                if (score >= 21) {
                    firstWins += count
                } else {
                    def newUniverse = new Universe(
                            firstPosition: position,
                            secondPosition: universe.secondPosition,
                            firstScore: score,
                            secondScore: universe.secondScore)
                    newUniverses[newUniverse] = newUniverses.getOrDefault(newUniverse, 0L) + count
                }
            }
        }
        universes = newUniverses

        if (!universes)
            return [firstWins, secondWins]

        println("Universes before second roll: ${universes.values().sum()}")

        newUniverses = [:]
        universes.each { universe, count ->
            splits().each { dice ->
                def (int score, int position) = roll(universe.secondPosition, universe.secondScore, dice)
                if (score >= 21) {
                    secondWins += count
                } else {
                    def newUniverse = new Universe(
                            firstPosition: universe.firstPosition,
                            secondPosition: position,
                            firstScore: universe.firstScore,
                            secondScore: score)
                    newUniverses[newUniverse] = newUniverses.getOrDefault(newUniverse, 0L) + count
                }
            }
        }

        simulate(newUniverses, firstWins, secondWins)
    }

    @Memoized
    static List<Integer> roll(int position, int score, int dice) {
        def newPosition = position + dice
        if (newPosition > 10) newPosition -= 10
        def newScore = score + newPosition
        [newScore, newPosition]
    }

    @Memoized
    static List<Byte> splits() {
        [1, 2, 3].collectMany {
            i ->
                [1, 2, 3].collectMany {
                    j -> [1, 2, 3].collect { k -> i + j + k }
                }
        }
    }

    static void main(String... args) {
        def (first, second) = simulate(4, 10)
        println(max(first as long, second as long))
    }
}
