package allergen

import (
	"errors"
	"fmt"
	"github.com/advendofcode/util"
	"sort"
	"strings"
)

type Food struct {
	ingredients []string
	allergens   []string
}

type Products struct {
	Safe      []string
	Allergens map[string]string
}

func solve(foods []Food) Products {
	products, err := removeNext(foods, Products{[]string{}, map[string]string{}})
	if err != nil {
		panic(err)
	}
	return products
}

func canonical(allergens map[string]string) string {
	var ingredients []string
	as := keys(allergens)
	sort.Strings(as)
	for _, allergen := range as {
		ingredients = append(ingredients, allergens[allergen])
	}

	return strings.Join(ingredients, ",")
}

func keys(m map[string]string) []string {
	var result []string
	for k := range m {
		result = append(result, k)
	}
	return result
}

func removeNext(previous []Food, acc Products) (Products, error) {
	sort.Slice(previous, byAllergensAndIngredients(previous))

	if len(previous) == 0 {
		return Products{[]string{}, map[string]string{}}, nil
	}

	var head Food
	head, err := findFirst(previous, func(food Food) bool {
		return len(food.allergens) > 0
	})
	if err != nil {
		for _, f := range previous {
			if len(f.allergens) == 0 {
				for _, i := range f.ingredients {
					acc.Safe = append(acc.Safe, i)
				}
			}
		}
		return acc, nil
	}
	if len(head.allergens) > 1 {
		panic("not implemented")
	}
	allergen := head.allergens[0]

ingredients:
	for _, ingredient := range head.ingredients {
		var next []Food
		for _, f := range previous {
			if contains(f.allergens, allergen) && !contains(f.ingredients, ingredient) {
				continue ingredients
			}

			allergens := remove(f.allergens, allergen)
			ingredients := remove(f.ingredients, ingredient)
			next = append(next, Food{ingredients, allergens})
		}
		products, err := removeNext(next, acc)
		if err != nil {
			continue
		}
		products.Allergens[allergen] = ingredient
		return products, nil
	}
	return Products{}, errors.New("failed to remove " + allergen)
}

func findFirst(foods []Food, f func(Food) bool) (Food, error) {
	for _, food := range foods {
		if f(food) {
			return food, nil
		}
	}
	return Food{}, errors.New("no matches")
}

func remove(xs []string, v string) []string {
	var result []string
	for _, x := range xs {
		if x != v {
			result = append(result, x)
		}
	}
	return result
}

func contains(xs []string, v string) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

func byAllergensAndIngredients(foods []Food) func(i int, j int) bool {
	return func(i, j int) bool {
		if len(foods[i].allergens) == len(foods[j].allergens) {
			return len(foods[i].ingredients) < len(foods[j].ingredients)
		}
		return len(foods[i].allergens) < len(foods[j].allergens)
	}
}

func parse(lines []string) []Food {
	var result []Food
	for _, line := range lines {
		products, allergens := util.Split(line, " \\(contains ")
		allergens = strings.ReplaceAll(allergens, ")", "")
		result = append(result, Food{
			strings.Split(products, " "),
			strings.Split(allergens, ", "),
		})
	}
	return result
}

func Solve() {
	file := "2020/allergen/puzzle.txt"
	foods := parse(util.Lines(file))
	ps := solve(foods)
	fmt.Println(len(ps.Safe))
	fmt.Println(canonical(ps.Allergens))
}
