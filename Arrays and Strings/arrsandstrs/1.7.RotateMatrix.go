package arrsandstrs

type row = []int
type matrix = []row

// RotateMatrix rotates a matrix by 90 degrees clockwise. The original matrix is left intact.
func RotateMatrix(m matrix) matrix {
	if len(m) == 0 {
		return m
	}
	var res []row
	for j := 0; j < len(m); j++ {
		if len(m[j]) == 0 {
			return m
		}
		var row row
		for i := len(m) - 1; i >= 0; i-- {
			row = append(row, m[i][j])
		}
		res = append(res, row)
	}
	return res
}

// RotateMatrix rotates a matrix by 90 degrees clockwise in place.
// WARN: The original matrix is changed.
// func RotateMatrixInPlace(m matrix) matrix {
// 	if len(m) == 0 {
// 		return m
// 	}
// 	for j := 0; j < len(m); j++ {
// 		if len(m[j]) == 0 {
// 			return m
// 		}
// 		tmp := -1
// 		for i, y := len(m)-1, 0; i >= 0; i, y = i-1, y+1 {
// 			tmp = m[y][j]
// 			m[y][j] = m[i][j]
// 			// println()
// 			println(i, j)
// 			println(y, j)
// 		}
// 	}
// 	return m
// }
