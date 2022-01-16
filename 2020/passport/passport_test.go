package passport

import (
	"gotest.tools/assert"
	"testing"
)

func TestSampleInput(t *testing.T) {
	// given:
	input := "example.txt"

	// when:
	result := solve(input)

	// then:
	assert.Equal(t, result, 2)
}

var invalidPassports = []string{
	"eyr:1972 cid:100\nhcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
	"iyr:2019\nhcl:#602927 eyr:1967 hgt:170cm\necl:grn pid:012533040 byr:1946",
	"hcl:dab227 iyr:2012\necl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
	"hgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007",
}

var validPassports = []string{
	"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\nhcl:#623a2f",
	"eyr:2029 ecl:blu cid:129 byr:1989\niyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
	"hcl:#888785\nhgt:164cm byr:2001 iyr:2015 cid:88\npid:545766238 ecl:hzl\neyr:2022",
	"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
}

var edgeCases = []struct {
	name  string
	data  string
	valid bool
}{
	{
		name:  "max iyr",
		data:  "iyr:2020 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
		valid: true,
	},
	{
		name:  "min iyr",
		data:  "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
		valid: true,
	},
	{
		name:  "max height in cm",
		data:  "iyr:2010 hgt:193cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
		valid: true,
	},
	{
		name:  "190 cm",
		data:  "iyr:2010 hgt:190cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
		valid: true,
	},
}

func TestVariousPassports(t *testing.T) {
	for _, data := range invalidPassports {
		t.Run(data, func(t *testing.T) {
			// expect:
			assert.Equal(t, valid(data), false)
		})
	}

	for _, data := range validPassports {
		t.Run(data, func(t *testing.T) {
			// expect:
			assert.Equal(t, valid(data), true)
		})
	}

	for _, tc := range edgeCases {
		t.Run(tc.name, func(t *testing.T) {
			// expect:
			assert.Equal(t, valid(tc.data), tc.valid)
		})
	}
}
