import static java.lang.Math.floorDiv

class SnailFish {
    static reduce(number) {
        while (true) {
            def exploded = explode(number)
            if (number != exploded) {
                number = exploded
                continue
            }
            def splitted = split(number)
            if (splitted != number) {
                number = splitted
                continue
            }
            return number
        }
    }

    static explode(number) {
        explode(number, 0)[0]
    }

    static explode(List number, int level) {
        def (left, right) = number
        if (left instanceof Integer && right instanceof Integer) {
            return level >= 4 ? [0, left, right] : [number, 0, 0]
        }
        if (left instanceof List) {
            def (exploded, l, r) = explode(left, level + 1)
            if (exploded != left)
                return [[exploded, addLeft(right, r)], l as int, 0]
        }
        if (right instanceof List) {
            def (exploded, l, r) = explode(right, level + 1)
            return [[addRight(left, l), exploded], 0, r as int]
        }
        return [number, 0, 0]
    }

    static addRight(number, int delta) {
        if (number instanceof Integer) return number + delta
        def (left, right) = number
        [left, addRight(right, delta)]
    }

    static addLeft(number, int delta) {
        if (number instanceof Integer) return number + delta
        def (left, right) = number
        [addLeft(left, delta), right]
    }

    static split(number) {
        if (number instanceof Integer) {
            if (number < 10) return number
            def left = floorDiv(number as int, 2)
            return [left, number - left]
        }
        def (left, right) = number
        def leftSplit = split(left)
        left == leftSplit ? [left, split(right)] : [split(left), right]
    }

    static add(List... xs) {
        add(xs.toList())
    }

    static add(List xs) {
        xs.inject { a, b -> reduce([a, b]) }
    }

    static int bestMagnitude(List... numbers) {
        bestMagnitude(numbers.toList())
    }

    static int bestMagnitude(List numbers) {
        numbers.collectMany { a ->
            a = a as List
            numbers.collectMany { b ->
                b = b as List
                a != b ? ([add(a, b), add(b, a)].collect { magnitude(it) }) : []
            }
        }.max()
    }

    static int magnitude(number) {
        if (number instanceof Integer) return number
        def (left, right) = number
        magnitude(left) * 3 + magnitude(right) * 2
    }

    static void main(String... args) {
        def input = [
                [[8, [6, [0, 1]]], [8, [2, [1, 9]]]],
                [[[5, [1, 7]], 9], [4, [[9, 4], 4]]],
                [[[1, [6, 3]], 8], [[3, [5, 3]], [1, 6]]],
                [0, [[1, [7, 4]], [6, [7, 3]]]],
                [[[7, 6], 2], [[[9, 5], 9], [[5, 5], [6, 3]]]],
                [[3, [[2, 3], 0]], 0],
                [4, [2, [4, [2, 6]]]],
                [[2, [8, [4, 5]]], [[[7, 7], 7], 1]],
                [[[[6, 7], 6], [9, 8]], [[0, [7, 3]], [[9, 1], [2, 0]]]],
                [0, [[[8, 4], [4, 9]], 5]],
                [[0, 7], [8, 5]],
                [6, [[9, [0, 7]], [[0, 0], [8, 1]]]],
                [[[[8, 3], [1, 9]], [[9, 3], [6, 5]]], 4],
                [[[6, 6], [[1, 2], [1, 7]]], [[8, 8], [3, 2]]],
                [[[6, 4], [[0, 3], 1]], [[6, 2], [4, [0, 3]]]],
                [[[2, 9], [[2, 1], 1]], [[6, [1, 4]], [6, [0, 3]]]],
                [9, [[7, 4], [1, 9]]],
                [[[[1, 2], [7, 8]], [[9, 6], [1, 3]]], [[0, 6], [[3, 6], 6]]],
                [[[1, [8, 6]], [2, [3, 4]]], [[0, 4], [5, [5, 7]]]],
                [[[5, 9], [[1, 0], [4, 3]]], [[7, [6, 7]], [1, [1, 5]]]],
                [[[[6, 8], [5, 9]], [5, [4, 5]]], 4],
                [[[[3, 4], 4], 7], [[5, [3, 7]], 7]],
                [[[[3, 3], [7, 9]], [1, [4, 8]]], 0],
                [[[3, [9, 4]], [1, [3, 7]]], [[[1, 8], 1], [6, 1]]],
                [3, [[[7, 5], [4, 8]], [7, 8]]],
                [[7, [2, [2, 4]]], [0, [8, [0, 3]]]],
                [[[[8, 5], 3], [3, [8, 3]]], [1, [0, [7, 4]]]],
                [[[[7, 1], 3], [3, 4]], [[3, 7], [[1, 8], [4, 8]]]],
                [[[3, [9, 9]], 4], [[4, 2], [[4, 2], 4]]],
                [[[5, [9, 1]], [[3, 5], [1, 9]]], 7],
                [[[[0, 8], 5], [9, [5, 1]]], [[7, 0], 1]],
                [[[0, 2], [[1, 9], 7]], [[0, 3], [[0, 3], [4, 8]]]],
                [[[[1, 8], 0], [[8, 6], [7, 6]]], [[[1, 8], 4], [[0, 4], [8, 3]]]],
                [[[1, [2, 7]], [[5, 4], [5, 0]]], [5, [8, [8, 4]]]],
                [[[[4, 4], [7, 3]], [4, [2, 3]]], [[[6, 5], [1, 5]], [5, [8, 6]]]],
                [[[[7, 8], 4], [9, [4, 2]]], [[[1, 4], 2], [0, 7]]],
                [[8, 4], [1, [2, 5]]],
                [[[[2, 5], 4], [7, [0, 2]]], [5, 3]],
                [[3, [[7, 4], 3]], 3],
                [[[3, 5], [3, [1, 4]]], [[[0, 8], 1], 8]],
                [[[[1, 9], 5], [2, [4, 8]]], [[[9, 2], [0, 1]], 1]],
                [[[6, [1, 5]], [[2, 2], 6]], [[1, [2, 6]], 5]],
                [[[3, 2], [9, 3]], [[2, 1], [4, 8]]],
                [[[[9, 2], 7], [[5, 9], [1, 2]]], [[[3, 0], [2, 8]], 0]],
                [[[6, 5], [[9, 4], 3]], [[[6, 2], 1], [0, 7]]],
                [[[8, 6], 1], [9, [1, [0, 1]]]],
                [[[[5, 1], 4], [8, [6, 8]]], [4, [[1, 8], 9]]],
                [[[[1, 1], [8, 9]], [2, [0, 6]]], 3],
                [[[1, [8, 3]], [[4, 3], 1]], [[[4, 1], [8, 6]], 8]],
                [[8, [[6, 2], 8]], [[[4, 0], 8], 6]],
                [[[[2, 2], 7], [[9, 0], [3, 3]]], [[[4, 4], 0], 2]],
                [8, [[3, [9, 1]], [0, [9, 1]]]],
                [[[0, [4, 2]], [[2, 2], [8, 7]]], [[6, [4, 2]], [1, 6]]],
                [[3, 2], [4, [[6, 2], 2]]],
                [[6, [3, [2, 9]]], [[9, [1, 5]], [4, 4]]],
                [[[[7, 5], 5], 8], [1, [0, [2, 7]]]],
                [[2, [[2, 9], [1, 6]]], [[[0, 1], [0, 2]], [4, [3, 4]]]],
                [[[[8, 9], [7, 4]], [8, [6, 5]]], 1],
                [[8, 9], [[2, [6, 9]], [2, 8]]],
                [[5, 1], 8],
                [[[8, [4, 2]], [5, [1, 8]]], [[0, [0, 6]], [[6, 7], 9]]],
                [[[8, [8, 0]], [[8, 0], 8]], [[[9, 9], 9], [9, [5, 4]]]],
                [[[[3, 3], 5], [5, [9, 0]]], [[2, 6], [[3, 8], [7, 1]]]],
                [3, [[[1, 5], 8], 5]],
                [[[9, 8], [4, 3]], 5],
                [[[5, 7], [[2, 1], 6]], [[4, 2], [1, [0, 2]]]],
                [[[[9, 3], [9, 8]], [[1, 0], 6]], [[[6, 5], 2], [[0, 3], 6]]],
                [8, [[[9, 8], [2, 8]], [1, 0]]],
                [[8, [5, 9]], [[[4, 3], 6], [[5, 1], 4]]],
                [[0, 8], [1, [4, [6, 3]]]],
                [3, [3, [6, [5, 6]]]],
                [0, [0, [[8, 0], 8]]],
                [[0, 4], [[7, 4], [[0, 7], 1]]],
                [7, [[[6, 3], [4, 0]], 1]],
                [9, [5, [[5, 3], [2, 8]]]],
                [[7, [[8, 3], [1, 7]]], [[[2, 7], 1], [[9, 4], [7, 1]]]],
                [[[0, [7, 3]], 3], 2],
                [[1, [[9, 0], 2]], 3],
                [[1, [7, [0, 1]]], [[1, 8], 5]],
                [3, [5, [4, 1]]],
                [3, [[[9, 8], 4], [4, [9, 7]]]],
                [[2, 9], [0, 9]],
                [[[[7, 1], [9, 3]], [1, [1, 8]]], 9],
                [[[9, 8], [[7, 8], 3]], [[1, [6, 3]], [2, [7, 3]]]],
                [[[7, 3], [1, [5, 5]]], [[4, 8], [8, [2, 5]]]],
                [[2, [[6, 5], [4, 6]]], [[0, 3], 7]],
                [[[4, [9, 7]], [[6, 1], 6]], [[[8, 1], 6], [[2, 5], 9]]],
                [[[6, 0], 0], [9, 9]],
                [[[[1, 0], 0], [[5, 7], 9]], [[[7, 2], 0], [9, 6]]],
                [[[[5, 0], [2, 0]], [0, [7, 5]]], [[[7, 7], [2, 4]], 8]],
                [0, [[9, [3, 4]], [[3, 4], 6]]],
                [[[0, 8], [[1, 5], [3, 4]]], [[5, [6, 4]], [[2, 5], [2, 5]]]],
                [[8, 0], [[2, [7, 9]], 9]],
                [[[3, [7, 0]], [3, [8, 4]]], 2],
                [[8, 1], [[[8, 9], [1, 0]], 3]],
                [[[8, 3], [[4, 8], 4]], [[8, [8, 8]], [0, 2]]],
                [[0, [9, 4]], [[6, 8], [[7, 1], 9]]],
                [[[[5, 3], [2, 8]], [8, 7]], [9, [[5, 9], [5, 2]]]],
                [2, [4, [[4, 3], 8]]],
                [[[[7, 2], [6, 4]], 7], 8]
        ]
        def solution = add(input)
        println(solution)
        println(magnitude(solution))
        println(bestMagnitude(input))
    }
}
