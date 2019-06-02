package easy

import (
	"strconv"
	"strings"
)

// You need to construct a string consists of parenthesis and integers
// from a binary tree with the preorder traversing way.
//
// The null node needs to be represented by empty parenthesis pair "()".
// And you need to omit all the empty parenthesis pairs that don't affect
// the one-to-one mapping relationship between the string and the original binary tree.
//
// Example 1:
// Input: Binary tree: [1,2,3,4]
//        1
//      /   \
//     2     3
//    /
//   4
//
// Output: "1(2(4))(3)"
//
// Explanation: Originallay it needs to be "1(2(4)())(3()())",
// but you need to omit all the unnecessary empty parenthesis pairs.
// And it will be "1(2(4))(3)".
//
// Example 2:
// Input: Binary tree: [1,2,3,null,4]
//        1
//      /   \
//     2     3
//      \
//       4
//
// Output: "1(2()(4))(3)"
//
// Explanation: Almost the same as the first example,
// except we can't omit the first parenthesis pair to break
// the one-to-one mapping relationship between the input and the output.

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	var str []string
	preorderStr(t, &str)

	s := strings.Join(str, "")
	s = strings.TrimPrefix(s, "(")
	s = strings.TrimSuffix(s, ")")
	return s
}

func preorderStr(t *TreeNode, str *[]string) {
	*str = append(*str, "(", strconv.Itoa(t.Val))

	if t.Left == nil && t.Right == nil {
		*str = append(*str, ")")
		return
	}

	if t.Left == nil {
		*str = append(*str, "()")
	} else {
		preorderStr(t.Left, str)
	}

	if t.Right != nil {
		preorderStr(t.Right, str)
	}
	*str = append(*str, ")")
}

// The official solution:
// /**
//  * Definition for a binary tree node.
//  * public class TreeNode {
//  *     int val;
//  *     TreeNode left;
//  *     TreeNode right;
//  *     TreeNode(int x) { val = x; }
//  * }
//  */
//
// Approach #1 Using Recursion:
//
//  public class Solution {
//     public String tree2str(TreeNode t) {
//         if(t==null)
//             return "";
//         if(t.left==null && t.right==null)
//             return t.val+"";
//         if(t.right==null)
//             return t.val+"("+tree2str(t.left)+")";
//         return t.val+"("+tree2str(t.left)+")("+tree2str(t.right)+")";
//     }
// }
//
// Approach #2 Iterative Method Using stack:
//
// public class Solution {
//     public String tree2str(TreeNode t) {
//         if (t == null)
//             return "";
//         Stack < TreeNode > stack = new Stack < > ();
//         stack.push(t);
//         Set < TreeNode > visited = new HashSet < > ();
//         StringBuilder s = new StringBuilder();
//         while (!stack.isEmpty()) {
//             t = stack.peek();
//             if (visited.contains(t)) {
//                 stack.pop();
//                 s.append(")");
//             } else {
//                 visited.add(t);
//                 s.append("(" + t.val);
//                 if (t.left == null && t.right != null)
//                     s.append("()");
//                 if (t.right != null)
//                     stack.push(t.right);
//                 if (t.left != null)
//                     stack.push(t.left);
//             }
//         }
//         return s.substring(1, s.length() - 1);
//     }
// }
