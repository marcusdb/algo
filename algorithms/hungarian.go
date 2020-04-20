package algorithms

import (
	"math"
)

func arrayCopy(matrix [][]float64) [][]float64 {
	n := len(matrix)
	m := len(matrix[0])
	duplicate := make([][]float64, n)
	data := make([]float64, n*m)
	for i := range matrix {
		start := i * m
		end := start + m
		duplicate[i] = data[start:end:end]
		copy(duplicate[i], matrix[i])
	}
	return duplicate

}

func zeroOutMatrix(input [][]float64) [][]float64 {
	for row := range input {
		for col := range input {
			input[row][col] = 0
		}
	}
	return input
}

// Solve receives an MxN matrix and returns a with the chosen row for each column.
// as per https://www.youtube.com/watch?v=rrfFTdO2Z7I
func Solve(input [][]float64) (float64, []int) {

	var total float64
	local := arrayCopy(input)
	if len(input) != len(input[0]) {
		local = normalizeMatrix(local)
	}

	result := make([]int, len(input[0]))
	if len(input) > len(input[0]) {
		//missing colunms
		local = zeroRows(local)
		local = zeroColumns(local)
	} else {
		// equals or missing rows
		local = zeroRows(local)
		local = zeroColumns(local)
	}

	marked := makeArray(len(input))
	var markedColumns map[int]bool
	var markedRows map[int]bool
	marked, markedColumns = rowScanning(local, marked)
	marked, markedRows = colScanning(local, marked)
	for checkLines(marked) != len(input) {

		local = step4(local, markedRows, markedColumns)
		marked = zeroOutMatrix(marked)
		marked, markedColumns = rowScanning(local, marked)
		marked, markedRows = colScanning(local, marked)
	}

	for row := range input {
		for col := range input[0] {
			if marked[row][col] == 1 {

				total += input[row][col]
				result[row] = col
			}

		}
	}
	return total, result
}

func makeArray(dimension int) [][]float64 {
	a := make([][]float64, dimension)
	for i := range a {
		a[i] = make([]float64, dimension)
	}
	return a
}

// checkLines check how many lines are covered
func checkLines(marked [][]float64) int {
	total := 0
	for row := range marked {
		for col := range marked {
			if marked[row][col] == 1 {
				total++
			}
		}
	}
	return total
}

func step4(input [][]float64, rows map[int]bool, columns map[int]bool) [][]float64 {
	local := input
	// calculate min of the not cover by lines
	min := math.MaxFloat64
	for row := range input {
		for col := range input {
			if !rows[row] && !columns[col] {
				if input[row][col] < min {
					min = input[row][col]
				}
			}
		}
	}

	for row := range input {
		for col := range input {
			if !rows[row] && !columns[col] {
				local[row][col] = input[row][col] - min
			}
			if rows[row] && columns[col] {
				local[row][col] = input[row][col] + min
			}
		}
	}
	return local
}

//colScanning return the marked array (1 assigned -1 cross out) and the crossed out rows
//col scan mark the col with one zeros and cross out its rows
func colScanning(input [][]float64, marked [][]float64) ([][]float64, map[int]bool) {
	markedRow := make(map[int]bool)
	for col := range input {
		chosenRow := -1
		for row := range input {
			if input[row][col] == 0 && marked[row][col] == 0 {
				if chosenRow >= 0 { // already had another zero
					chosenRow = -1
					break
				}
				chosenRow = row
			}
		}
		if chosenRow >= 0 {
			marked[chosenRow][col] = 1
			markedRow[chosenRow] = true
			for col2 := range input {

				if input[chosenRow][col2] == 0 && col != col2 {
					marked[chosenRow][col2] = -1
				}

			}
		}
	}
	return marked, markedRow
}

//rowScanning return the marked array (1 assigned -1 cross out) and the crossed out columns
//row scan mark the row with one zeros and cross out its colums
func rowScanning(input [][]float64, marked [][]float64) ([][]float64, map[int]bool) {
	markedCol := make(map[int]bool)
	for row := range input {
		chosenCol := -1
		for col := range input {
			if input[row][col] == 0 && marked[row][col] == 0 {
				if chosenCol >= 0 { // already had another zero
					chosenCol = -1
					break
				}
				chosenCol = col
			}
		}
		if chosenCol >= 0 {
			marked[row][chosenCol] = 1
			markedCol[chosenCol] = true
			for row2 := range input {

				if input[row2][chosenCol] == 0 && row != row2 {
					marked[row2][chosenCol] = -1
				}

			}
		}
	}
	return marked, markedCol
}

// cover all the zeros with the minimum possible number of lines
func coverZeros(input [][]float64) int {

	return 0
}
func calculateNumberOfLines(markedColumns map[int]bool, markedRows map[int]bool, dimension int) int {
	result := len(markedColumns)
	for i := 0; i < dimension; i++ {
		if !markedRows[i] {
			result++
		}
	}
	return result
}

func zeroRows(input [][]float64) [][]float64 {
	localMatrix := input
	minValues := make([]float64, len(localMatrix))

	// find min for each row
	for col := range localMatrix[0] { // go through columns
		for row := range localMatrix { // for each row
			if col == 0 { // first column
				minValues[row] = localMatrix[row][0]
			}
			if minValues[row] > localMatrix[row][col] {
				minValues[row] = localMatrix[row][col]
			}
		}
	}

	// now subtract all values of the row by the min
	for row := range localMatrix {
		for col := range localMatrix[row] { // go through columns
			localMatrix[row][col] = localMatrix[row][col] - minValues[row]
		}
	}
	return localMatrix
}

func zeroColumns(input [][]float64) [][]float64 {
	localMatrix := input
	minValues := make([]float64, len(localMatrix))

	// find min for each column
	for col := range localMatrix[0] { // go through columns
		for row := range localMatrix { // for each row
			if row == 0 { // first column
				minValues[col] = localMatrix[0][col]
			}
			if minValues[col] > localMatrix[row][col] {
				minValues[col] = localMatrix[row][col]
			}
		}
	}

	// now subtract all values of the column by the min
	for row := range localMatrix {
		for col := range localMatrix[row] { // go through columns
			localMatrix[row][col] = localMatrix[row][col] - minValues[col]
		}
	}
	return localMatrix
}

//normalize receives and MxN matrix and makes it max(m,n) matrix
func normalizeMatrix(originalMatrix [][]float64) [][]float64 {

	// assuming there are missing columns
	maxDimension := len(originalMatrix) //total rows
	if maxDimension < len(originalMatrix[0]) {
		//there are missing rows
		maxDimension = len(originalMatrix[0]) // total columns
	}
	localMatrix := makeArray(maxDimension)
	var maxValue float64 = 0
	for row := range originalMatrix { // go through rows
		for col := range originalMatrix[row] { // go through columns
			if maxValue < originalMatrix[row][col] {
				maxValue = originalMatrix[row][col]
			}
		}
	}

	// fill as many as the missing columns with rows max values
	for i := 0; i < maxDimension; i++ {
		for j := 0; j < maxDimension; j++ {
			if i >= len(originalMatrix) || j >= len(originalMatrix[0]) {
				localMatrix[i][j] = maxValue
			} else {
				localMatrix[i][j] = originalMatrix[i][j]
			}

		}
	}
	return localMatrix

}
