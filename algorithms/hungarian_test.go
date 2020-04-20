package algorithms

import "testing"

var testSolve = []struct {
	name   string
	data   [][]float64
	result []int
	total  float64
}{
	{
		name: "test 1",
		data: [][]float64{
			{9, 11, 14, 11, 7},
			{6, 15, 13, 13, 10},
			{12, 13, 6, 8, 8},
			{11, 9, 10, 12, 9},
			{7, 12, 14, 10, 14},
		},
		result: []int{4, 0, 2, 1, 3},
		total:  38,
	},
	{
		name: "test 2",
		data: [][]float64{
			{108, 125, 150},
			{150, 135, 175},
			{122, 148, 250},
		},
		result: []int{2, 1, 0},
		total:  407,
	},
	/*	{
			name: "test 3",
			data: [][]float64{
				{108, 125, 150, 260},
				{150, 135, 175, 260},
				{122, 148, 250, 260},
				{134, 200, 260, 260},
			},
			result: []int{2, 1, 0},
			total:  407,
		},
	*/
	{
		name: "test 4",
		data: [][]float64{
			{11, 6, 12},
			{12, 4, 6},
			{8, 12, 11},
		},
		result: []int{1, 2, 0},
		total:  20,
	},
	{
		name: "test 5",
		data: [][]float64{
			{13, 13, 19, 50, 33, 38},
			{73, 33, 71, 77, 97, 95},
			{20, 8, 56, 55, 64, 35},
			{26, 25, 72, 32, 55, 77},
			{83, 40, 69, 3, 53, 49},
			{67, 20, 44, 29, 86, 61},
		},
		result: []int{4, 1, 5, 0, 3, 2},
		total:  174,
	},
	{
		name: "test 6 with missing column",
		data: [][]float64{
			{11, 6, 12},
			{12, 4, 6},
			{8, 12, 11},
			{14, 16, 15},
		},
		result: []int{1, 2, 0},
		total:  20,
	},
	// {
	// 	name: "test 7 with missing row",
	// 	data: [][]float64{
	// 		{11, 6, 12, 15},
	// 		{12, 4, 6, 17},
	// 		{8, 12, 11, 18},
	// 	},
	// 	result: []int{1, 2, 0},
	// 	total:  20,
	// },
}

var testStep4 = []struct {
	name    string
	data    [][]float64
	rows    map[int]bool
	columns map[int]bool
	result  [][]float64
}{
	{
		name: "test 1",
		data: [][]float64{
			{2, 4, 7, 2, 0},
			{0, 9, 7, 5, 4},
			{6, 7, 0, 0, 2},
			{2, 0, 1, 1, 0},
			{0, 5, 7, 1, 7},
		},
		rows:    map[int]bool{2: true},
		columns: map[int]bool{0: true, 1: true, 4: true},
		result: [][]float64{
			{2, 4, 6, 1, 0},
			{0, 9, 6, 4, 4},
			{7, 8, 0, 0, 3},
			{2, 0, 0, 0, 0},
			{0, 5, 6, 0, 7},
		},
	},
}

var testColScanning = []struct {
	name   string
	data   [][]float64
	marked [][]float64
	result [][]float64
}{
	{
		name: "test 1",
		data: [][]float64{
			{2, 4, 7, 2, 0},
			{0, 9, 7, 5, 4},
			{6, 7, 0, 0, 2},
			{2, 0, 1, 1, 0},
			{0, 5, 7, 1, 7},
		},
		marked: [][]float64{
			{0, 0, 0, 0, 1},
			{1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, -1},
			{-1, 0, 0, 0, 0},
		},
		result: [][]float64{
			{0, 0, 0, 0, 1},
			{1, 0, 0, 0, 0},
			{0, 0, 1, -1, 0},
			{0, 1, 0, 0, -1},
			{-1, 0, 0, 0, 0},
		}},
}

var testRowScanning = []struct {
	name   string
	data   [][]float64
	marked [][]float64
	result [][]float64
}{
	{
		name: "test 1",
		data: [][]float64{
			{2, 4, 7, 2, 0},
			{0, 9, 7, 5, 4},
			{6, 7, 0, 0, 2},
			{2, 0, 1, 1, 0},
			{0, 5, 7, 1, 7},
		},
		marked: [][]float64{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
		result: [][]float64{
			{0, 0, 0, 0, 1},
			{1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, -1},
			{-1, 0, 0, 0, 0},
		}},
}

var testsNormalize = []struct {
	name   string
	data   [][]float64
	result [][]float64
}{
	{
		name: "missing 1 row",
		data: [][]float64{
			{6, 2, 3, 4, 5},
			{3, 8, 2, 8, 1},
			{9, 9, 5, 4, 2},
			{6, 7, 3, 4, 3},
		},
		result: [][]float64{
			{6, 2, 3, 4, 5},
			{3, 8, 2, 8, 1},
			{9, 9, 5, 4, 2},
			{6, 7, 3, 4, 3},
			{9, 9, 9, 9, 9},
		}},
	{name: "missing 1 column",
		data: [][]float64{
			{6, 2, 3, 4},
			{3, 8, 2, 8},
			{9, 9, 5, 4},
			{6, 7, 3, 4},
			{9, 9, 5, 8},
		},
		result: [][]float64{
			{6, 2, 3, 4, 9},
			{3, 8, 2, 8, 9},
			{9, 9, 5, 4, 9},
			{6, 7, 3, 4, 9},
			{9, 9, 5, 8, 9},
		}},
	{name: "missing 2 columns",
		data: [][]float64{
			{6, 2, 3},
			{3, 8, 2},
			{9, 9, 5},
			{6, 7, 3},
			{9, 9, 5},
		},
		result: [][]float64{
			{6, 2, 3, 9, 9},
			{3, 8, 2, 9, 9},
			{9, 9, 5, 9, 9},
			{6, 7, 3, 9, 9},
			{9, 9, 5, 9, 9},
		}},
}

var testsZeroRows = []struct {
	name   string
	data   [][]float64
	result [][]float64
}{
	{
		name: "test 1",
		data: [][]float64{
			{6, 2, 3, 4, 5},
			{3, 8, 2, 8, 1},
			{9, 9, 5, 4, 2},
			{6, 7, 3, 4, 3},
			{9, 9, 5, 8, 5},
		},
		result: [][]float64{
			{4, 0, 1, 2, 3},
			{2, 7, 1, 7, 0},
			{7, 7, 3, 2, 0},
			{3, 4, 0, 1, 0},
			{4, 4, 0, 3, 0},
		},
	},
}

var testsZeroCols = []struct {
	name   string
	data   [][]float64
	result [][]float64
}{
	{
		name: "test 1",
		data: [][]float64{
			{6, 2, 3, 4, 5},
			{3, 8, 2, 8, 1},
			{9, 9, 5, 4, 2},
			{6, 7, 3, 4, 3},
			{9, 9, 5, 8, 5},
		},
		result: [][]float64{
			{3, 0, 1, 0, 4},
			{0, 6, 0, 4, 0},
			{6, 7, 3, 0, 1},
			{3, 5, 1, 0, 2},
			{6, 7, 3, 4, 4},
		},
	},
}

func TestCheckLines(t *testing.T) {
	data := [][]float64{
		{0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 1, -1, 0},
		{0, 1, 0, 0, -1},
		{-1, 0, 0, 0, 0},
	}

	result := checkLines(data)
	if result != 4 {
		t.Errorf("TestCheckLines want %v got %v", 4, result)
	}

}

func TestStep4(t *testing.T) {
	for _, testToBeDone := range testStep4 {
		result := step4(testToBeDone.data, testToBeDone.rows, testToBeDone.columns)
		if !compareMulti(result, testToBeDone.result) {
			t.Errorf("%s want %v got %v", testToBeDone.name, testToBeDone.result, result)
		}
	}
}

func TestNormalizeMatrix(t *testing.T) {
	for _, testToBeDone := range testsNormalize {
		result := normalizeMatrix(testToBeDone.data)
		if !compareMulti(result, testToBeDone.result) {
			t.Errorf("%s want %v got %v", testToBeDone.name, testToBeDone.result, result)
		}
	}
}

func TestRowScanning(t *testing.T) {
	for _, testToBeDone := range testRowScanning {
		result, _ := rowScanning(testToBeDone.data, testToBeDone.marked)
		if !compareMulti(result, testToBeDone.result) {
			t.Errorf("%s want %v got %v", testToBeDone.name, testToBeDone.result, result)
		}
	}
}

func TestSolve(t *testing.T) {
	for _, testToBeDone := range testSolve {
		total, columns := Solve(testToBeDone.data)
		if !compareSimple(columns, testToBeDone.result) {
			t.Errorf("%s want %v got %v", testToBeDone.name, testToBeDone.result, columns)
		}
		if total != testToBeDone.total {
			t.Errorf("%s want %v got %v", testToBeDone.name, testToBeDone.total, total)
		}
	}
}

func TestColScanning(t *testing.T) {
	for _, testToBeDone := range testColScanning {
		result, _ := colScanning(testToBeDone.data, testToBeDone.marked)
		if !compareMulti(result, testToBeDone.result) {
			t.Errorf("%s want %v got %v", testToBeDone.name, testToBeDone.result, result)
		}
	}
}

func TestZeroRows(t *testing.T) {
	for _, testToBeDone := range testsZeroRows {
		result := zeroRows(testToBeDone.data)
		if !compareMulti(result, testToBeDone.result) {
			t.Errorf("%s want %v got %v", testToBeDone.name, testToBeDone.result, result)
		}
	}
}

func TestZeroColumns(t *testing.T) {
	for _, testToBeDone := range testsZeroCols {
		result := zeroColumns(testToBeDone.data)
		if !compareMulti(result, testToBeDone.result) {
			t.Errorf("%s want %v got %v", testToBeDone.name, testToBeDone.result, result)
		}
	}
}
func compareSimple(m []int, n []int) bool {
	if len(m) != len(n) {
		return false
	}
	for i := range m {
		if m[i] != n[i] {
			return false
		}
	}
	return true
}

func compareMulti(m [][]float64, n [][]float64) bool {
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
