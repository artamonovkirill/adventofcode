package cups

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestParses(t *testing.T) {
	cups := parse("389125467")
	expected := []string{
		"{3, previous: 7, next: 8, previousByValue: 2}",
		"{8, previous: 3, next: 9, previousByValue: 7}",
		"{9, previous: 8, next: 1, previousByValue: 8}",
		"{1, previous: 9, next: 2, previousByValue: 9}",
		"{2, previous: 1, next: 5, previousByValue: 1}",
		"{5, previous: 2, next: 4, previousByValue: 4}",
		"{4, previous: 5, next: 6, previousByValue: 3}",
		"{6, previous: 4, next: 7, previousByValue: 5}",
		"{7, previous: 6, next: 3, previousByValue: 6}",
	}

	current := cups.current

	for i := 0; i < cups.len; i++ {
		fmt.Println(current.ToString())
		assert.Equal(t, current.ToString(), expected[i])
		current = current.next
	}

	assert.Equal(t, cups.ToString(), "25467389")
}

func TestAdvancesExample(t *testing.T) {
	assert.Equal(t, advance("389125467", 10), "92658374")
	assert.Equal(t, advance("389125467", 100), "67384529")
}

func TestSolvesExample(t *testing.T) {
	input := "389125467"
	assert.Equal(t, solve(input), 149245887792)
}
