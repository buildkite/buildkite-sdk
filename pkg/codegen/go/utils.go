package go_code_gen

var goUtilsFile = `// This file is auto-generated please do not edit

package buildkite

import "strconv"

func ParseStringToInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return val
}

func ParseStringSliceToIntSlice(strs []string) []int {
	var result []int
	for _, str := range strs {
		val, err := strconv.Atoi(str)
		if err != nil {
			result = append(result, 0)
			continue
		}
		result = append(result, val)
	}

	return result
}

func ParseStringToBool(str string) bool {
	val, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return val
}

func ParseStringSliceToBoolSlice(strs []string) []bool {
	var result []bool
	for _, str := range strs {
		val, err := strconv.ParseBool(str)
		if err != nil {
			result = append(result, false)
			continue
		}
		result = append(result, val)
	}

	return result
}`

func newUtilsFiles() string {
	return goUtilsFile
}
