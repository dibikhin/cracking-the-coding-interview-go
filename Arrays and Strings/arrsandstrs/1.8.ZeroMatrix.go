package arrsandstrs

import "fmt"

// ZeroMatrix sets entire row and column to 0 if an element is 0. The original matrix left intact
func ZeroMatrix(m matrix) matrix {
	res := make(matrix, 0, len(m))
	for i := 0; i < len(m); i++ {
		row := make(row, len(m))
		res = append(res, row)
		for j := 0; j < len(m); j++ {
			// v := m[i][j]
			// if v != 0 {
			// 	row = append(row, v)
			// }
			// println(m[i][j])
		}
	}
	fmt.Println(res)
	return res
}
