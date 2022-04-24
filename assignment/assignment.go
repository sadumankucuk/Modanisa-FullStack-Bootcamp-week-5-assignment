package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	//uint32 is the set of all unsigned 32-bit integers. Range: 0 through 4294967295.
	sum64 := uint64(x) + uint64(y)
	overflow := uint32(sum64>>32) == 1
	return x + y, overflow
}

func CeilNumber(f float64) float64 {
	unit := 0.25
	return math.Ceil(f/unit) * unit
}

func AlphabetSoup(s string) string {
	result := strings.Split(s, "")
	sort.Strings(result)
	return strings.Join(result, "")
}

func StringMask(s string, n uint) string {
	if len(s) == 0 {
		return "*"
	} else if len(s) <= int(n) {
		return strings.Repeat("*", len(s))
	} else {
		return s[:n] + strings.Repeat("*", len(s)-int(n))
	}
}

// For Benchmark Test
func StringMaskReyyanSolve(s string, n uint) string {
	stringArray := strings.Split(s, "")
	if len(stringArray) <= 1 {
		return "*"
	}
	if int(n) >= len(stringArray) {
		n = 0
	}
	for j := int(n); j < len(stringArray); j++ {
		stringArray[j] = "*"
	}
	justString := strings.Join(stringArray, "")
	return justString
}

func WordSplit(arr [2]string) string {
	words := strings.Split(arr[1], ",")
	s := strings.Split(arr[0], "")

	var findWord []string
	var results []string

	for _, letter := range s {
		findWord = append(findWord, letter)
		for _, word := range words {
			if word == strings.Join(findWord, "") {
				findWord = []string{}
				results = append(results, word)
			}
		}
	}

	if len(strings.Join(results, ",")) != 0 {
		return strings.Join(results, ", ")
	} else {
		return "not possible"
	}
}

func VariadicSet(i ...interface{}) []interface{} {
	//The make function allocates a zeroed array and returns a slice that refers to that array
	keys := make(map[interface{}]bool)
	var list []interface{}
	for _, entry := range i {
		if value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
