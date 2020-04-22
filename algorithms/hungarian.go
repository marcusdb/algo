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

	marked = rowScanning(local, marked)

	lineRows, lineCols := checkLines(marked)

	for len(lineRows)+len(lineCols) != len(input) {
		local = step4(local, lineRows, lineCols)
		marked = zeroOutMatrix(marked)
		marked = rowScanning(local, marked)
		lineRows, lineCols = checkLines(marked)
		break
	}
	//fmt.Printf("local:%v\nr:%v\nc:%v\nmarked:%v\n", local, lineRows, lineCols, marked)

	result = resultDecider(marked)[:len(input[0])]

	for row := range input {
		total += input[row][result[row]]
	}
	return total, result
}

func resultDecider(input [][]float64) []int {
	resultPerRow := make(map[int]int)
	// check if there is only 1 possible per row and marked
	for row := range input {
		chosenCol := -1
		total := 0
		for col := range input {
			if input[row][col] != 0 {
				total++
				chosenCol = col
			}
		}
		if total == 1 {
			resultPerRow[row] = chosenCol
		}
	}
	//fmt.Printf("row alone %v\n", resultPerRow)
	if len(resultPerRow) < len(input) {
		// check if there is only 1 possible per col
		for col := range input {
			chosenRow := -1
			total := 0
			for row := range input {
				if input[row][col] != 0 {
					total++
					chosenRow = row
				}
			}
			if total == 1 {
				resultPerRow[chosenRow] = col
			}
		}
		//fmt.Printf("col alone %v\n", resultPerRow)
		if len(resultPerRow) < len(input) {
			for row := range input {
				for col := range input {
					if _, ok := resultPerRow[row]; !ok {
						if input[row][col] != 0 {
							// check for no use of this column
							notUsed := true
							for key := range resultPerRow {
								if resultPerRow[key] == col {
									notUsed = false
								}
							}
							if notUsed {
								resultPerRow[row] = col
							}
						}
					}
				}
			}
		}
	}
	result := make([]int, len(input))
	for key := range resultPerRow {
		result[key] = resultPerRow[key]
	}
	return result
}

func makeArray(dimension int) [][]float64 {
	a := make([][]float64, dimension)
	for i := range a {
		a[i] = make([]float64, dimension)
	}
	return a
}

// checkLines check how many lines are covered
func checkLines(marked [][]float64) (map[int]bool, map[int]bool) {
	markedRows := make(map[int]bool)
	markedCols := make(map[int]bool)

	//Mark all rows having no assignments
	for row := range marked {
		markedRows[row] = true
		for col := range marked {
			if marked[row][col] == 1 {
				delete(markedRows, row)
			}
		}
	}

	var numMarkedRows int
	var numMarkedCols int
	newlyMarkedRows := markedRows
	newlyMarkedCols := make(map[int]bool)
	for numMarkedRows != len(markedRows) || numMarkedCols != len(markedCols) {
		numMarkedRows = len(markedRows)
		numMarkedCols = len(markedCols)
		//Mark all columns having zeros in newly marked row(s)
		for markedRow := range newlyMarkedRows {
			for col := range marked {
				if marked[markedRow][col] != 0 {
					newlyMarkedCols[col] = true
					markedCols[col] = true
				}
			}
		}
		newlyMarkedRows = make(map[int]bool) // reset newly
		//Mark all rows having assignments in newly marked columns
		for markedCol := range newlyMarkedCols {
			for row := range marked {
				if marked[row][markedCol] == 1 {
					newlyMarkedRows[row] = true
					markedRows[row] = true
				}
			}
		}
		newlyMarkedCols = make(map[int]bool) // reset newly

	}
	// return lines for the markedCols and for the UNmarkedRows
	rowLines := make(map[int]bool)
	for i := range marked {
		if !markedRows[i] {
			rowLines[i] = true
		}
	}
	colLines := markedCols
	return rowLines, colLines
}

func step4(input [][]float64, rows map[int]bool, columns map[int]bool) [][]float64 {
	local := arrayCopy(input)
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

//rowScanning return the marked array (1 assigned -1 cross out) and the crossed out columns
//row scan mark the row with one zeros and cross out its colums
func rowScanning(input [][]float64, marked [][]float64) [][]float64 {
	for row := range input {

		for col := range input {
			if input[row][col] == 0 && marked[row][col] == 0 {

				marked[row][col] = 1
				for i := col; i < len(input); i++ {
					if input[row][i] == 0 && marked[row][i] == 0 {
						// cross out all left in the row
						marked[row][i] = -1
					}
				}
				for row2 := range input {
					// cross all in the same column
					if input[row2][col] == 0 && row != row2 {
						marked[row2][col] = -1
					}

				}
				break
			}
		}
	}
	return marked
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
