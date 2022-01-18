package combo

import "fmt"

func loopSize(subject int, public int) int {
	loop := 1
	for {
		if key(subject, loop) == public {
			return loop
		}
		loop++
	}
}

func key(subject int, loop int) int {
	result := 1
	for i := 0; i < loop; i++ {
		result *= subject
		result %= 20201227
	}
	return result
}

func Solve() {
	loop := loopSize(7, 9659666)
	fmt.Println(key(loop, 75188))
}
