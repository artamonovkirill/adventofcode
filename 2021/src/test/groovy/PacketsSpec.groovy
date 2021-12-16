import spock.lang.Specification

class PacketsSpec extends Specification {

    def 'decodes literal packet'() {
        given:
        def binary = Packets.binary('D2FE28')

        when:
        def (packet, rest) = Packets.packet(binary)

        then:
        packet.version == 6
        packet.value == 2021
        rest =~ /0+/
    }

    def 'decodes operator packet'() {
        given:
        def binary = Packets.binary(hex)

        when:
        def (packet, rest) = Packets.packet(binary)

        then:
        packet.version == version
        packet.packets*.value == values
        rest =~ /0*/

        where:
        hex              | version | values
        '38006F45291200' | 1       | [10, 20]
        'EE00D40C823060' | 7       | [1, 2, 3]
    }

    def 'decodes nested packet'() {
        given:
        def binary = Packets.binary(hex)

        when:
        def (packet, rest) = Packets.packet(binary)

        then:
        packet.versions().sum() == sum
        rest =~ /0*/

        where:
        hex                              | sum
        '8A004A801A8002F478'             | 16
        '620080001611562C8802118E34'     | 12
        'C0015000016115A2E0802F182340'   | 23
        'A0016C880162017C3686B18A3D4780' | 31
    }

    def 'calculates result'() {
        given:
        def binary = Packets.binary(hex)

        when:
        def (packet, rest) = Packets.packet(binary)

        then:
        packet.calculate() == expected
        rest =~ /0*/

        where:
        hex                          | expected
        'C200B40A82'                 | 3
        '04005AC33890'               | 54
        '880086C3E88112'             | 7
        'CE00C43D881120'             | 9
        'D8005AC2A8F0'               | 1
        'F600BC2D8F'                 | 0
        '9C005AC2F8F0'               | 0
        '9C0141080250320F1802104A08' | 1
    }

}
