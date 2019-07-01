package easy

// In an infinite binary tree where every node has two children,
// the nodes are labelled in row order.
//
// In the odd numbered rows (ie., the first, third, fifth,...),
// the labelling is left to right, while in the even numbered
// rows (second, fourth, sixth,...), the labelling is right to left.
//
// Image location: [/image/easy/path_in_zigzag_labelled_binary_tree.png]
//
// Given the label of a node in this tree, return the labels in the path
// from the root of the tree to the node with that label.
//
// Example 1:
// Input: label = 14
// Output: [1,3,4,14]
//
// Example 2:
// Input: label = 26
// Output: [1,2,6,10,26]
//
// Constraints:
// 1 <= label <= 10^6

func pathInZigZagTree(label int) []int {
	if label <= 0 {
		return nil
	}

	labels := []int{0, 1} // lables[0] is sentinel
	totalNodeNum, upLevelNodeNum, round := 1, 1, 2
	for totalNodeNum < label {
		newCap := upLevelNodeNum * 2
		labels = append(labels, make([]int, newCap)...)
		if round%2 == 0 {
			for i := newCap; i < 2*newCap; i++ {
				labels[i] = 3*newCap - i - 1
			}
		} else {
			for i := newCap; i < 2*newCap; i++ {
				labels[i] = i
			}
		}

		totalNodeNum += newCap
		upLevelNodeNum = newCap
		round++
	}

	idx := label / 2
	for labels[idx] != label {
		idx++
	}

	var path []int
	for labels[idx] != 0 {
		path = append(path, labels[idx])
		idx = idx / 2
	}
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
	}

	return path
}
