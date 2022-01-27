package hard

func maximumGood(statements [][]int) int {
	max := 0 // maximum good person
	n := len(statements)
	for i := 0; i < (1 << n); i++ { // for every possible state set
		count := 0 // count good person numer
	Tag:
		for j := 0; j < n; j++ { // prepare to find every good person
			if (1<<j)&i == 0 { // j is a bad person
				continue
			}

			count++
			for l, v := range statements[j] {
				// // NOTE: for debug
				// if i == 7 {
				// 	fmt.Println(l, v, j, count, (1<<l)&i)
				// }

				if v == 0 && ((1<<l)&i) != 0 { // good person say other good person is a bad person
					count = 0
					break Tag
				}
				if v == 1 && ((1<<l)&i) == 0 { // good person say other bad person is a good person
					count = 0
					break Tag
				}
			}
		}
		if count > max {
			max = count
		}

	}

	return max
}
