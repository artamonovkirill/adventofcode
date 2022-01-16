package cups

import (
	"fmt"
	"github.com/advendofcode/util"
	"strconv"
	"strings"
)

type Node struct {
	value           int
	previous        *Node
	next            *Node
	previousByValue *Node
}

func (n *Node) ToString() string {
	var next int
	if n.next != nil {
		next = n.next.value
	}
	return fmt.Sprintf("{%d, previous: %d, next: %d, previousByValue: %d}",
		n.value, n.previous.value, next, n.previousByValue.value)
}

type Cups struct {
	current *Node
	one     *Node
	len     int
	max     int
}

func number(c uint8) int {
	return util.Number(string(c))
}

func parse(input string) *Cups {
	first := &Node{number(input[0]), nil, nil, nil}
	last := first
	var one *Node
	var max int
	nodes := map[int]*Node{
		first.value: first,
	}
	for i := 1; i < len(input); i++ {
		value := number(input[i])
		if value > max {
			max = value
		}
		n := &Node{value, last, nil, nil}
		last.next = n
		nodes[n.value] = n
		last = n
		if value == 1 {
			one = n
		}
	}
	first.previous = last
	last.next = first

	for _, node := range nodes {
		var previous *Node
		if node.value > 1 {
			previous = nodes[node.value-1]
		} else {
			previous = nodes[max]
		}
		node.previousByValue = previous
	}

	return &Cups{first, one, len(input), max}
}

func (c *Cups) maxNode() *Node {
	current := c.current
	for {
		if current.value == c.max {
			return current
		}
		current = current.next
	}
}

func extend(cups *Cups) {
	extendTo := 1_000_000
	max := cups.maxNode()
	first := &Node{max.value + 1, cups.current.previous, nil, max}
	cups.current.previous.next = first
	last := first
	for i := cups.max + 2; i <= extendTo; i++ {
		n := &Node{i, last, nil, last}
		last.next = n
		last = n
	}
	last.previous = cups.current.previous
	last.next = cups.current
	cups.current.previous = last
	cups.max = extendTo
	cups.len = extendTo
	cups.one.previousByValue = last
}

func solve(input string) int {
	reportRate := 1_000_000
	cups := parse(input)
	extend(cups)
	for i := 0; i < 10_000_000; i++ {
		if i%reportRate == 0 {
			fmt.Printf("%d0%%\n", i/reportRate)
		}
		move(cups)
	}
	return cups.one.next.value * cups.one.next.next.value
}

func advance(input string, moves int) string {
	cups := parse(input)
	for i := 0; i < moves; i++ {
		move(cups)
	}
	return cups.ToString()
}

func move(cups *Cups) {
	pick1 := cups.current.next
	pick2 := pick1.next
	pick3 := pick2.next

	cups.current.next = pick3.next
	pick3.next.previous = cups.current

	destination := cups.current.previousByValue
	for destination.value == pick1.value || destination.value == pick2.value || destination.value == pick3.value {
		destination = destination.previousByValue
	}

	pick1.previous = destination
	pick3.next = destination.next
	destination.next.previous = pick3
	destination.next = pick1
	cups.current = cups.current.next
}

func (c *Cups) ToString() string {
	values := make([]string, c.len-1)
	current := c.one
	for i := 0; i < c.len-1; i++ {
		current = current.next
		values[i] = strconv.Itoa(current.value)
	}
	return strings.Join(values, "")
}

func Solve() {
	fmt.Println(advance("589174263", 10000000))
	fmt.Println(solve("589174263"))
}
