import groovy.transform.EqualsAndHashCode

import static java.lang.Integer.parseInt
import static java.lang.Long.parseLong

abstract class Packet {
    int version

    List<Integer> versions() {
        [version]
    }

    abstract long calculate()
}

@EqualsAndHashCode
class Literal extends Packet {
    long value

    @Override
    String toString() {
        "L($version): $value"
    }

    @Override
    long calculate() {
        value
    }
}

@EqualsAndHashCode
class Operation extends Packet {
    List<Packet> packets
    int type

    @Override
    String toString() {
        "$version: $packets"
    }

    @Override
    List<Integer> versions() {
        [version] + packets.collectMany { it.versions() }
    }

    @Override
    long calculate() {
        switch (type) {
            case 0: return packets*.calculate().sum() as long
            case 1: return packets*.calculate().inject(1 as long) { a, b -> a * b }
            case 2: return packets*.calculate().min()
            case 3: return packets*.calculate().max()
            case 5:
                def (a, b) = packets*.calculate()
                return a > b ? 1 : 0
            case 6:
                def (a, b) = packets*.calculate()
                return a < b ? 1 : 0
            case 7:
                def (a, b) = packets*.calculate()
                return a == b ? 1 : 0
            default: throw new RuntimeException('not implemented')
        }
    }
}

class Bits {
    String value

    Bits(String value) {
        this.value = value
    }

    String next(int length) {
        def result = value.substring(0, length)
        value = value.substring(length)
        return result
    }
}

class Packets {
    private static final LITERAL = 4
    private static final Map<String, String> DIGITS = [
            '0': '0000',
            '1': '0001',
            '2': '0010',
            '3': '0011',
            '4': '0100',
            '5': '0101',
            '6': '0110',
            '7': '0111',
            '8': '1000',
            '9': '1001',
            'A': '1010',
            'B': '1011',
            'C': '1100',
            'D': '1101',
            'E': '1110',
            'F': '1111']

    static String binary(String hex) {
        hex.toList().collect { DIGITS[it] }.join('')
    }

    static packet(String binary) {
        def bits = new Bits(binary)
        def version = parseInt(bits.next(3), 2)
        def type = parseInt(bits.next(3), 2)
        if (type == LITERAL) {
            def match = bits.value.find(/(1[0-1]{4})*0[0-1]{4}/)
            bits.next(match.length())
            def value = match.toList().collate(5)*.tail()*.join('').join('')
            return [new Literal(version: version, value: parseLong(value, 2)), bits.value]
        } else {
            def length = bits.next(1) == '0' ? 15 : 11
            def subpackets = parseInt(bits.next(length), 2)
            def packets = []
            if (length == 15) {
                def rest = bits.next(subpackets)
                while (rest) {
                    def (value, r) = packet(rest)
                    packets << value
                    rest = r
                }
                return [new Operation(version: version, type: type, packets: packets), bits.value]
            } else {
                def rest = bits.value
                while (subpackets > 0) {
                    def (value, r) = packet(rest)
                    packets << value
                    rest = r
                    subpackets--
                }
                return [new Operation(version: version, type: type, packets: packets), rest]
            }
        }
    }

    static void main(String... args) {
        def hex = '420D598021E0084A07C98EC91DCAE0B880287912A925799429825980593D7DCD400820329480BF21003CC0086028910097520230C80813401D8CC00F601881805705003CC00E200E98400F50031801D160048E5AFEFD5E5C02B93F2F4C11CADBBB799CB294C5FDB8E12C40139B7C98AFA8B2600DCBAF4D3A4C27CB54EA6F5390B1004B93E2F40097CA2ECF70C1001F296EF9A647F5BFC48C012C0090E675DF644A675DF645A7E6FE600BE004872B1B4AAB5273ED601D2CD240145F802F2CFD31EFBD4D64DD802738333992F9FFE69CAF088C010E0040A5CC65CD25774830A80372F9D78FA4F56CB6CDDC148034E9B8D2F189FD002AF3918AECD23100953600900021D1863142400043214C668CB31F073005A6E467600BCB1F4B1D2805930092F99C69C6292409CE6C4A4F530F100365E8CC600ACCDB75F8A50025F2361C9D248EF25B662014870035600042A1DC77890200D41086B0FE4E918D82CC015C00DCC0010F8FF112358002150DE194529E9F7B9EE064C015B005C401B8470F60C080371460CC469BA7091802F39BE6252858720AC2098B596D40208A53CBF3594092FF7B41B3004A5DB25C864A37EF82C401C9BCFE94B7EBE2D961892E0C1006A32C4160094CDF53E1E4CDF53E1D8005FD3B8B7642D3B4EB9C4D819194C0159F1ED00526B38ACF6D73915F3005EC0179C359E129EFDEFEEF1950005988E001C9C799ABCE39588BB2DA86EB9ACA22840191C8DFBE1DC005EE55167EFF89510010B322925A7F85A40194680252885238D7374C457A6830C012965AE00D4C40188B306E3580021319239C2298C4ED288A1802B1AF001A298FD53E63F54B7004A68B25A94BEBAAA00276980330CE0942620042E3944289A600DC388351BDC00C9DCDCFC8050E00043E2AC788EE200EC2088919C0010A82F0922710040F289B28E524632AE0'
        def binary = binary(hex)
        def (packet, rest) = Packets.packet(binary)
        assert rest =~ /0*/
        println(packet.calculate())
    }
}
