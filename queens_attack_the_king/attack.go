package queens_attack_the_king

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	queensMap := make(map[int]map[int]bool)
	for _, q := range queens {
		r, c := q[0], q[1]
		if _, ok := queensMap[r]; !ok {
			queensMap[r] = map[int]bool{}
		}
		queensMap[r][c] = true
	}

	directions := [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
		[]int{1, -1},
		[]int{1, 1},
		[]int{-1, 1},
		[]int{-1, -1},
	}

	var attackQueens [][]int
	for _, d := range directions {
		x, y := d[0], d[1]
		for r, c := king[0], king[1]; r < 8 && r >= 0 && c < 8 && c >= 0; r, c = r+x, c+y {
			if _, ok := queensMap[r][c]; ok {
				attackQueens = append(attackQueens, []int{r, c})
				break
			}
		}
	}

	return attackQueens
}
