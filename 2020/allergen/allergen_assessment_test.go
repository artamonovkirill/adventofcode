package allergen

import (
	"fmt"
	"github.com/advendofcode/util"
	"gotest.tools/assert"
	"testing"
)

func TestExample(t *testing.T) {
	// given:
	file := "example.txt"
	foods := parse(util.Lines(file))

	// when:
	products := solve(foods)

	// then:
	assert.Equal(t, len(products.Safe), 5)
	assert.Equal(t, contains(products.Safe, "kfcds"), true)
	assert.Equal(t, contains(products.Safe, "nhms"), true)
	assert.Equal(t, contains(products.Safe, "sbzzf"), true)
	assert.Equal(t, contains(products.Safe, "trh"), true)
}

func TestSolves(t *testing.T) {
	cases := []struct {
		foods    []Food
		products Products
	}{
		{[]Food{
			{[]string{"a"}, []string{"x"}},
		}, Products{[]string{}, map[string]string{"x": "a"}}},
		{[]Food{
			{[]string{"a", "b"}, []string{"x"}},
		}, Products{[]string{"b"}, map[string]string{"x": "a"}}},
		{[]Food{
			{[]string{"a", "b"}, []string{"x"}},
			{[]string{"a"}, []string{"y"}},
		}, Products{[]string{}, map[string]string{"y": "a", "x": "b"}}},
	}
	for _, c := range cases {
		t.Run(fmt.Sprint(c.foods), func(t *testing.T) {
			assert.DeepEqual(t, solve(c.foods), c.products)
		})
	}
}
