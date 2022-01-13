package jigsaw

import (
	"fmt"
	"github.com/advendofcode/util"
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestAssemblesExample(t *testing.T) {
	input := "example.txt"
	assert.Equal(t, corners(input), 20899048083289)
}

func TestRotates(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"rotations/0.txt", "rotations/90.txt"},
		{"rotations/90.txt", "rotations/180.txt"},
		{"rotations/180.txt", "rotations/270.txt"},
	}
	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			// given:
			input := parsePixels(util.Lines(c.input))
			expected := pixelsToString(parsePixels(util.Lines(c.expected)))

			// when:
			rotated := input.rotate()

			// then:
			assert.Equal(t, pixelsToString(rotated), expected)
		})
	}
}

func TestFlips(t *testing.T) {
	cases := []struct {
		input    string
		f        func(image Pixels) Pixels
		expected string
	}{
		{"flips/init.txt", flipHorizontally, "flips/horizontal.txt"},
		{"flips/init.txt", flipVertically, "flips/vertical.txt"},
	}
	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			// given:
			input := parsePixels(util.Lines(c.input))
			expected := pixelsToString(parsePixels(util.Lines(c.expected)))

			// when:
			rotated := c.f(input)

			// then:
			assert.Equal(t, pixelsToString(rotated), expected)
		})
	}
}

func TestAligns(t *testing.T) {
	cases := []struct {
		a        string
		b        string
		expected []Direction
	}{
		{"edges/center.txt", "edges/left.txt", []Direction{Left}},
		{"edges/center.txt", "edges/right.txt", []Direction{Right}},
		{"edges/center.txt", "edges/up.txt", []Direction{Up}},
		{"edges/center.txt", "edges/down.txt", []Direction{Down}},
		{"edges/dummy.txt", "edges/dummy.txt", []Direction{Left, Right, Up, Down}},
	}
	for _, c := range cases {
		t.Run(c.a, func(t *testing.T) {
			// given:
			from := parsePixels(util.Lines(c.a))
			to := parsePixels(util.Lines(c.b))

			// expect:
			assert.DeepEqual(t, from.align(to), c.expected)
		})
	}
}

func TestGeneratesExampleInnerVariations(t *testing.T) {
	cases := []string{
		"edges/center.txt",
		"edges/down.txt",
		"edges/left.txt",
		"edges/right.txt",
		"edges/up.txt",
	}
	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			images := parse("example.txt")
			expected := parsePixels(util.Lines(c))
			assertContains(t, images, expected)
		})
	}
}

func TestGeneratesExampleCornersVariations(t *testing.T) {
	cases := []struct {
		file string
		id   int
	}{
		{"example/00.txt", 1951},
		{"example/02.txt", 3079},
		{"example/20.txt", 2971},
		{"example/22.txt", 1171},
	}
	for _, c := range cases {
		t.Run(c.file, func(t *testing.T) {
			images := parse("example.txt")
			expected := parsePixels(util.Lines(c.file))

			contains := false
			for _, version := range images[c.id] {
				if reflect.DeepEqual(version, expected) {
					contains = true
					break
				}
			}
			assert.Equal(t, contains, true)
		})
	}
}

func assertContains(t *testing.T, images map[int][]Pixels, expected Pixels) {
	for _, versions := range images {
		for _, version := range versions {
			if reflect.DeepEqual(version, expected) {
				return
			}
		}
	}
	assert.Assert(t, false, "No match found")
}

func TestIntersect(t *testing.T) {
	cases := []struct {
		a        []Direction
		b        []Direction
		expected []Direction
	}{
		{[]Direction{Up}, []Direction{Down}, nil},
		{[]Direction{Down, Right}, []Direction{Down, Left}, []Direction{Down}},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v-%v-%v", c.a, c.b, c.expected), func(t *testing.T) {
			assert.DeepEqual(t, intersect(c.a, c.b), c.expected)
		})
	}
}

func TestFree(t *testing.T) {
	cases := []struct {
		i        Image
		p        Point
		expected []Direction
	}{
		{
			Image{0: map[int]Tile{0: {1, 1}}},
			Point{0, 0},
			[]Direction{Up, Down, Left, Right},
		},
		{
			Image{0: map[int]Tile{
				-1: {1, 1},
				0:  {1, 1},
				1:  {1, 1},
			}},
			Point{0, 0},
			[]Direction{Up, Down},
		},
		{
			Image{
				-1: map[int]Tile{
					0: {1, 1},
				},
				0: map[int]Tile{
					-1: {1, 1},
					0:  {1, 1},
					1:  {1, 1},
				},
				1: map[int]Tile{
					0: {1, 1},
				},
			},
			Point{0, 0},
			nil,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v-%v", c.i, c.p), func(t *testing.T) {
			assert.DeepEqual(t, free(c.i, c.p), c.expected)
		})
	}
}

func TestNeighbours(t *testing.T) {
	cases := []struct {
		i        Image
		p        Point
		expected map[Direction]Tile
	}{
		{
			Image{0: map[int]Tile{0: {1, 1}}},
			Point{0, 0},
			map[Direction]Tile{},
		},
		{
			Image{0: map[int]Tile{
				-1: {1, 1},
				0:  {2, 2},
				1:  {3, 3},
			}},
			Point{0, 0},
			map[Direction]Tile{
				Left:  {ID: 1, I: 1},
				Right: {ID: 3, I: 3},
			},
		},
		{
			Image{
				-1: map[int]Tile{
					0: {1, 1},
				},
				0: map[int]Tile{
					-1: {2, 2},
					0:  {3, 3},
					1:  {4, 4},
				},
				1: map[int]Tile{
					0: {5, 5},
				},
			},
			Point{0, 0},
			map[Direction]Tile{
				Up:    {ID: 1, I: 1},
				Left:  {ID: 2, I: 2},
				Right: {ID: 4, I: 4},
				Down:  {ID: 5, I: 5},
			},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v-%v", c.i, c.p), func(t *testing.T) {
			assert.DeepEqual(t, neighbours(c.p, c.i), c.expected)
		})
	}
}

func TestAddsOneTile(t *testing.T) {
	tiles := parse("example.txt")
	solution := parseSolution(tiles)
	fmt.Println(solution)
	delete(solution[0], 0)

	// when:
	_, err := arrange(solution, []int{1951}, tiles)

	// then:
	assert.NilError(t, err)
}

func parseSolution(tiles map[int][]Pixels) Image {
	result := make(Image, 3)
	for y := 0; y < 3; y++ {
		result[y] = make(map[int]Tile, 3)
		for x := 0; x < 3; x++ {
			file := fmt.Sprintf("example/%d%d.txt", y, x)
			lines := util.Lines(file)
			pixels := parsePixels(lines)
			for id, versions := range tiles {
				for i, version := range versions {
					if reflect.DeepEqual(pixels, version) {
						result[y][x] = Tile{id, i}
					}
				}
			}
		}
	}
	return result
}

func TestSolutionIsComplete(t *testing.T) {
	tiles := parse("example.txt")
	solution := parseSolution(tiles)

	image := join(solution, tiles)

	assert.Equal(t, pixelsToString(image), util.Text("example/image.txt"))
}

func TestJoinsImages(t *testing.T) {
	tiles := parse("example.txt")
	ids := keys(tiles)
	head := ids[0]
	init := Image{0: map[int]Tile{0: {head, 0}}}
	parts, err := arrange(init, remove(ids, head), tiles)

	assert.NilError(t, err)

	image := join(parts, tiles)

	assert.Equal(t, len(image.value), len(image.value[0]))

	matched := false
	for _, a := range alternate(image) {
		if reflect.DeepEqual(a, parsePixels(util.Lines("example/image.txt"))) {
			matched = true
		}
	}
	assert.Equal(t, matched, true)
}

func TestFindExampleMonsters(t *testing.T) {
	assert.Equal(t, monsterFree("example.txt", "monster.txt"), 273)
}

func TestRotatesMonster(t *testing.T) {
	original := parsePixels(util.Lines("monster.txt"))
	expected := parsePixels(util.Lines("rotations/m90.txt"))

	assert.Equal(t, pixelsToString(original.rotate()), pixelsToString(expected))
}
