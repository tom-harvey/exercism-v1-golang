// Package matrix can operate on Matrices
package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

// New reads numbers from a string and returns a rectanglar matrix of ints.
// Rows are separated by newlines; columns are separated by other whitespace.
func New(s string) (Matrix, error) {
	var ncol int
	var err error
	rows := strings.Split(s, "\n")
	m := make(Matrix, len(rows), len(rows))
	for i, r := range rows {
		columns := strings.Fields(r)
		if ncol == 0 { // first row
			if ncol = len(columns); ncol == 0 {
				err = fmt.Errorf("matrix.New: zero-length row")
				goto quit
			}
		} else { // subsequent row, require same length
			if c := len(columns); c != ncol {
				err = fmt.Errorf("row %d has %d columns, needs %d", i, c, ncol)
				goto quit
			}
		}
		m[i] = make([]int, ncol, ncol)
		for j, f := range columns {
			if m[i][j], err = strconv.Atoi(f); err != nil {
				goto quit
			}
		}
	}
quit:
	return m, err
}

// Dim returns the dimensions of m, or (0,0) if either is zero
func (m Matrix) Dim() (int, int) {
	var rows, columns int
	if rows = len(m); rows > 0 {
		if columns = len(m[0]); columns == 0 {
			rows = 0
		}
	}
	return rows, columns
}

// Rows returns a distinct copy of m (as a [][]int!)
func (m Matrix) Rows() [][]int {
	n := make(Matrix, len(m), len(m))
	for i, row := range m {
		n[i] = make([]int, len(row), len(row))
		copy(n[i], row)
	}
	return n
}

// Cols returns a distinct copy of the transpose of m (as a [][]int!)
func (m Matrix) Cols() [][]int {
	var n Matrix
	rows, columns := m.Dim()
	n = make(Matrix, columns, columns)
	for c := range n {
		n[c] = make([]int, rows, rows)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			n.Set(c, r, m[r][c])
		}
	}
	return n
}

// Set sets a value in m at the given row and column
func (m Matrix) Set(r, c, val int) bool {
	result := false
	if 0 <= r && r < len(m) && 0 <= c && c < len(m[r]) {
		m[r][c] = val
		result = true
	}
	return result
}
