package algorithms

import "testing"

var testsNormalize = []struct {
	name   string
	data   [][]int32
	result [][]int32
}{
	{
		name: "missing 1 row",
		data: [][]int32{
			{6, 2, 3, 4, 5},
			{3, 8, 2, 8, 1},
			{9, 9, 5, 4, 2},
			{6, 7, 3, 4, 3},
		},
		result: [][]int32{
			{6, 2, 3, 4, 5},
			{3, 8, 2, 8, 1},
			{9, 9, 5, 4, 2},
			{6, 7, 3, 4, 3},
			{9, 9, 5, 8, 5},
		}},
	{name: "missing 1 column",
		data: [][]int32{
			{6, 2, 3, 4},
			{3, 8, 2, 8},
			{9, 9, 5, 4},
			{6, 7, 3, 4},
			{9, 9, 5, 8},
		},
		result: [][]int32{
			{6, 2, 3, 4, 6},
			{3, 8, 2, 8, 8},
			{9, 9, 5, 4, 9},
			{6, 7, 3, 4, 7},
			{9, 9, 5, 8, 9},
		}},
	{name: "missing 2 columns",
		data: [][]int32{
			{6, 2, 3},
			{3, 8, 2},
			{9, 9, 5},
			{6, 7, 3},
			{9, 9, 5},
		},
		result: [][]int32{
			{6, 2, 3, 6, 6},
			{3, 8, 2, 8, 8},
			{9, 9, 5, 9, 9},
			{6, 7, 3, 7, 7},
			{9, 9, 5, 9, 9},
		}},
}

func TestNormalizeMatrix(t *testing.T) {
	for _, testToBeDone := range testsNormalize {
		result := normalizeMatrix(testToBeDone.data)
		if !compareMulti(result, testToBeDone.result) {
			t.Errorf("%s want %v got %v", testToBeDone.name, testToBeDone.result, result)
		}
	}

}

func compareMulti(m [][]int32, n [][]int32) bool {
	if len(m) != len(n) || len(m[0]) != len(n[0]) {
		return false
	}
	for i := range m {
		if len(m[i]) != len(n[i]) {
			return false
		}
		for j := range m[i] {
			if m[i][j] != n[i][j] {
				return false
			}
		}
	}
	return true
}
