import spock.lang.Specification

class SnailFishSpec extends Specification {
    def 'does not explode simple numbers'() {
        expect:
        SnailFish.explode(input) == output

        where:
        input                 | output
        [1, 2]                | [1, 2]
        [[[[0, 9], 2], 3], 4] | [[[[0, 9], 2], 3], 4]
    }

    def 'explodes to the right'() {
        expect:
        SnailFish.explode([[[[[9, 8], 1], 2], 3], 4]) == [[[[0, 9], 2], 3], 4]
    }

    def 'explodes to the left'() {
        expect:
        SnailFish.explode([7, [6, [5, [4, [3, 2]]]]]) == [7, [6, [5, [7, 0]]]]
    }

    def 'explodes to both sides'() {
        expect:
        SnailFish.explode([[6, [5, [4, [3, 2]]]], 1]) == [[6, [5, [7, 0]]], 3]
        SnailFish.explode([[3, [2, [1, [7, 3]]]], [6, [5, [4, [3, 2]]]]]) == [[3, [2, [8, 0]]], [9, [5, [4, [3, 2]]]]]
        SnailFish.explode([[3, [2, [8, 0]]], [9, [5, [4, [3, 2]]]]]) == [[3, [2, [8, 0]]], [9, [5, [7, 0]]]]
    }

    def 'splits'() {
        expect:
        SnailFish.split(5) == 5
        SnailFish.split(10) == [5, 5]
        SnailFish.split(11) == [5, 6]
        SnailFish.split(12) == [6, 6]
        SnailFish.split([[[[0, 7], 4], [15, [0, 13]]], [1, 1]]) == [[[[0, 7], 4], [[7, 8], [0, 13]]], [1, 1]]
        SnailFish.split([[[[0, 7], 4], [[7, 8], [0, 13]]], [1, 1]]) == [[[[0, 7], 4], [[7, 8], [0, [6, 7]]]], [1, 1]]
    }

    def 'reduces'() {
        expect:
        SnailFish.reduce(input) == output

        where:
        input                                          | output
        [1, 2]                                         | [1, 2]
        [[[[0, 9], 2], 3], 4]                          | [[[[0, 9], 2], 3], 4]
        [[[[[9, 8], 1], 2], 3], 4]                     | [[[[0, 9], 2], 3], 4]
        [7, [6, [5, [4, [3, 2]]]]]                     | [7, [6, [5, [7, 0]]]]
        [[6, [5, [4, [3, 2]]]], 1]                     | [[6, [5, [7, 0]]], 3]
        [[3, [2, [1, [7, 3]]]], [6, [5, [4, [3, 2]]]]] | [[3, [2, [8, 0]]], [9, [5, [7, 0]]]]
        [[[[[4, 3], 4], 4], [7, [[8, 4], 9]]], [1, 1]] | [[[[0, 7], 4], [[7, 8], [6, 0]]], [8, 1]]
    }

    def 'adds'() {
        expect:
        SnailFish.add([[[[4, 3], 4], 4], [7, [[8, 4], 9]]], [1, 1]) == [[[[0, 7], 4], [[7, 8], [6, 0]]], [8, 1]]
        SnailFish.add(
                [1, 1],
                [2, 2],
                [3, 3],
                [4, 4]) == [[[[1, 1], [2, 2]], [3, 3]], [4, 4]]
        SnailFish.add(
                [1, 1],
                [2, 2],
                [3, 3],
                [4, 4],
                [5, 5]) == [[[[3, 0], [5, 3]], [4, 4]], [5, 5]]
        SnailFish.add(
                [1, 1],
                [2, 2],
                [3, 3],
                [4, 4],
                [5, 5],
                [6, 6]) == [[[[5, 0], [7, 4]], [5, 5]], [6, 6]]
        SnailFish.add(
                [[[0, [4, 5]], [0, 0]], [[[4, 5], [2, 6]], [9, 5]]],
                [7, [[[3, 7], [4, 3]], [[6, 3], [8, 8]]]],
                [[2, [[0, 8], [3, 4]]], [[[6, 7], 1], [7, [1, 6]]]],
                [[[[2, 4], 7], [6, [0, 5]]], [[[6, 8], [2, 8]], [[2, 1], [4, 5]]]],
                [7, [5, [[3, 8], [1, 4]]]],
                [[2, [2, 2]], [8, [8, 1]]],
                [2, 9],
                [1, [[[9, 3], 9], [[9, 0], [0, 7]]]],
                [[[5, [7, 4]], 7], 1],
                [[[[4, 2], 2], 6], [8, 7]]
        ) == [[[[8, 7], [7, 7]], [[8, 6], [7, 7]]], [[[0, 7], [6, 6]], [8, 7]]]
    }

    def 'calculates magnitude'() {
        expect:
        SnailFish.magnitude([9, 1]) == 29
        SnailFish.magnitude([1, 9]) == 21
        SnailFish.magnitude([[9, 1], [1, 9]]) == 129
        SnailFish.magnitude([[1, 2], [[3, 4], 5]]) == 143
        SnailFish.magnitude([[[[0, 7], 4], [[7, 8], [6, 0]]], [8, 1]]) == 1384
        SnailFish.magnitude([[[[1, 1], [2, 2]], [3, 3]], [4, 4]]) == 445
        SnailFish.magnitude([[[[3, 0], [5, 3]], [4, 4]], [5, 5]]) == 791
        SnailFish.magnitude([[[[5, 0], [7, 4]], [5, 5]], [6, 6]]) == 1137
        SnailFish.magnitude([[[[8, 7], [7, 7]], [[8, 6], [7, 7]]], [[[0, 7], [6, 6]], [8, 7]]]) == 3488
    }

    def 'solves homework'() {
        when:
        def solution = SnailFish.add(
                [[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]],
                [[[5,[2,8]],4],[5,[[9,9],0]]],
                [6,[[[6,2],[5,6]],[[7,6],[4,7]]]],
                [[[6,[0,7]],[0,9]],[4,[9,[9,0]]]],
                [[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]],
                [[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]],
                [[[[5,4],[7,7]],8],[[8,3],8]],
                [[9,3],[[9,9],[6,[4,9]]]],
                [[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]],
                [[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]])

        then:
        solution == [[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]
        SnailFish.magnitude(solution) == 4140
    }

    def 'finds best magnitude'() {
        expect:
        SnailFish.bestMagnitude(
                [[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]],
                [[[5,[2,8]],4],[5,[[9,9],0]]],
                [6,[[[6,2],[5,6]],[[7,6],[4,7]]]],
                [[[6,[0,7]],[0,9]],[4,[9,[9,0]]]],
                [[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]],
                [[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]],
                [[[[5,4],[7,7]],8],[[8,3],8]],
                [[9,3],[[9,9],[6,[4,9]]]],
                [[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]],
                [[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]) == 3993
    }
}
