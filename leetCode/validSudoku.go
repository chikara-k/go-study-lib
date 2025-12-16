package leetcode

func isValidSudoku(board [][]byte) bool {
    rows := make([]map[byte]bool, 9)
    cols := make([]map[byte]bool, 9)
    boxes := make([]map[byte]bool, 9)

    for i := 0; i < 9; i++ {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        boxes[i] = make(map[byte]bool)
    }

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            v := board[r][c]
            if v == '.' {
                continue
            }

            box := (r/3)*3 + (c/3)

            if rows[r][v] {
                return false
            }

            if cols[c][v] {
                return false
            }

            if boxes[box][v] {
                return false
            }

            rows[r][v] = true
            cols[c][v] = true
            boxes[box][v] = true
        }
    }

    return true
}

// Use two-dimensional array
/*
func isValidSudoku(board [][]byte) bool {
	var rows [9][9]bool
	var cols [9][9]bool
	var boxes [9][9]bool

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			v := board[r][c]
			if v == '.' {
				continue
			}

			d := v - '1'              // '1'→0, ... '9'→8
			b := (r/3)*3 + (c/3)      // box index 0..8

			if rows[r][d] || cols[c][d] || boxes[b][d] {
				return false
			}

			rows[r][d] = true
			cols[c][d] = true
			boxes[b][d] = true
		}
	}

	return true
}
*/