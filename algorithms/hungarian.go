package algorithms

// Solve receives an MxN matrix and returns a with the chosen row for each column.

//normalize receives and MxN matrix and makes it max(m,n) matrix
func normalizeMatrix(originalMatrix [][]int32) [][]int32 {
	localMatrix := originalMatrix
	// assuming there are missing columns
	maxDimension := len(originalMatrix) //total rows

	if maxDimension < len(originalMatrix[0]) {
		//there are missing rows
		maxDimension = len(originalMatrix[0]) // total columns
		maxValues := make([]int32, maxDimension)

		// find maxValue per column
		for i := range originalMatrix { // go through rows
			for j := range originalMatrix[i] { // go through columns
				if i == 0 { // first line
					maxValues[j] = originalMatrix[0][j]
				}
				if maxValues[j] < originalMatrix[i][j] {
					maxValues[j] = originalMatrix[i][j]
				}
			}
		}
		// fill as many as the missing rows with columns max values
		for i := len(originalMatrix); i < maxDimension; i++ {
			localMatrix = append(localMatrix, maxValues)
		}
		return localMatrix
	}
	maxValues := make([]int32, maxDimension)
	// there are missing columns

	for col := range originalMatrix[0] { // go through columns
		for row := range originalMatrix { // for each row
			if col == 0 { // first column
				maxValues[row] = originalMatrix[row][0]
			}
			if maxValues[row] < originalMatrix[row][col] {
				maxValues[row] = originalMatrix[row][col]
			}
		}

	}
	// fill as many as the missing columns with rows max values
	for i := len(originalMatrix[0]); i < maxDimension; i++ {
		for row := range localMatrix {
			localMatrix[row] = append(localMatrix[row], maxValues[row])
		}

	}
	return localMatrix

}
