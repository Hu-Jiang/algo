package medium

// Given an integer n, generate all structurally unique
// BST's (binary search trees) that store values 1 ... n.
//
// Example:
// Input: 3
// Output:
// [
//   [1,null,3,2],
//   [3,2,null,1],
//   [3,1,null,null,2],
//   [2,1,3],
//   [1,null,2,null,3]
// ]
// Explanation:
// The above output corresponds to the 5 unique BST's shown below:
//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3

func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return nil
	}

	return doGenerateTrees(1, n)
}

func doGenerateTrees(left, right int) []*TreeNode {
	var trees []*TreeNode
	if left > right {
		trees = append(trees, nil)
		return trees
	}

	// chose i as root
	for i := left; i <= right; i++ {
		// all possible left subtrees if i is choosen to be a root
		leftSubTrees := doGenerateTrees(left, i-1)

		// all possible right subtrees if i is choosen to be a root
		rightSubTrees := doGenerateTrees(i+1, right)

		// connect left and right subtrees to the root i
		for j := 0; j < len(leftSubTrees); j++ {
			for k := 0; k < len(rightSubTrees); k++ {
				root := &TreeNode{i, leftSubTrees[j], rightSubTrees[k]}
				trees = append(trees, root)
			}
		}
	}

	return trees
}
