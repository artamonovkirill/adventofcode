import groovy.transform.ToString

@ToString
class Position {
    int x = 0
    int depth = 0
    int aim = 0
}

class Dive {
    static Position execute(List<String> commands) {
        commands.inject(new Position()) {
            p, c -> execute(p, c)
        }
    }

    static Position execute(Position p, String command) {
        def parts = command.split(' ')
        def directive = parts[0]
        def value = parts[1] as int
        switch (directive) {
            case "up":
                return new Position(x: p.x, depth: p.depth, aim: p.aim - value)
            case "down":
                return new Position(x: p.x, depth: p.depth, aim: p.aim + value)
            default:
                return new Position(x: p.x + value, depth: p.depth + p.aim * value, aim: p.aim)
        }
    }

    static void main(String... args) {
        def input = new File('src/main/resources/dive.txt')
        def commands = input.readLines()
        def result = execute(commands)
        System.out.println(result.x * result.depth)
    }
}
