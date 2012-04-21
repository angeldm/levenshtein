package levenshtein

func LevenshteinDistance(a, b string) int {
	n := len(a)
	m := len(b)
	if n == 0 || m == 0 {
		return -1
	}

	max := 0
	maxs := ""
	min := 0
	mins := ""

	if n > m {
		max = n
		min = m
		maxs = a
		mins = b
	} else {
		max = m
		min = n
		maxs = b
		mins = a

	}

	mat := make([][]int, max+1)
	for i := range mat {
		mat[i] = make([]int, min+1)
		mat[i][0] = i
	}
	for i := range mat[0] {
		mat[0][i] = i
	}

	for i := 1; i <= max; i++ {
		for j := 1; j <= min; j++ {
			d := 1
			if maxs[i-1] == mins[j-1] {
				d = 0
			}
			mat[i][j] = minimum(mat[i-1][j]+1, mat[i][j-1]+1, mat[i-1][j-1]+d)

		}
	}

	return mat[max][min]
}

func minimum(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

