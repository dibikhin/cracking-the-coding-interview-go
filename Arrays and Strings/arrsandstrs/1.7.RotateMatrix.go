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
