package distress

import (
	"gotest.tools/assert"
	"testing"
)

func TestSolvesExample(t *testing.T) {
	// expect:
	assert.Equal(t, count("example.txt"), 13)
	assert.Equal(t, arrange("example.txt"), 140)
}

func TestParses(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected []interface{}
	}{
		{"[1,2,3]", []interface{}{1, 2, 3}},
		{"[[1]]", []interface{}{[]interface{}{1}}},
		{"[[1],[2,3]]", []interface{}{[]interface{}{1}, []interface{}{2, 3}}},
		{"[[1],4]", []interface{}{[]interface{}{1}, 4}},
		{"[[],[10,5,6]]", []interface{}{[]interface{}{}, []interface{}{10, 5, 6}}},
		{"[[],[9,[5,6]]]", []interface{}{[]interface{}{}, []interface{}{9, []interface{}{5, 6}}}},
	} {
		t.Run(tc.input, func(t *testing.T) {
			/// when:
			result := parse(tc.input)

			// expect:
			assert.DeepEqual(t, result, tc.expected)
		})
	}
}

func TestCompares(t *testing.T) {
	for _, tc := range []struct {
		a        string
		b        string
		expected int
	}{
		{"[[],[8,[],5],[]]", "[[9]]", right},
		{"[[[[],[1,2]],[[],[9,5,5],7,8],0,3],[4,[3,1],0]]", "[[10,[[8,3,9],10]]]", right},
		{"[[8,2,5,5,7]]", "[[],[10,5,6],[7,[],0],[9,[[7,9],[5,4,10],[9,8],1,3],7,10,[]]]", wrong},
	} {
		t.Run(tc.a+"vs. "+tc.b, func(t *testing.T) {
			// given:
			a := parse(tc.a)
			b := parse(tc.b)

			// when:
			result := compare(a, b)

			// expect:
			assert.DeepEqual(t, result, tc.expected)
		})
	}
}
