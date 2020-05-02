//
// mix of simple array examples
//

package main

import (
	"fmt"
)

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func traverseRecursive(root *treeNode, path string) []string {
	var ltree []string
	var rtree []string
	var res []string

	if root == nil {
		return []string{}
	}

	ltree = traverseRecursive(root.Left, path+"L")
	res = append(res, ltree...)

	res = append(res, fmt.Sprintf("%d:%s", root.Val, path))

	rtree = traverseRecursive(root.Right, path+"R")
	res = append(res, rtree...)

	return res
}

func main() {

	// Create tree:
	//           0
	//          / \
	//         1   3
	//        /     \
	//       2       4
	//              / \
	//             5   6
	//

	root := treeNode{0, nil, nil}

	l := treeNode{1, nil, nil}
	r := treeNode{3, nil, nil}

	ll := treeNode{2, nil, nil}
	rr := treeNode{4, nil, nil}

	rrl := treeNode{5, nil, nil}
	rrr := treeNode{6, nil, nil}

	root.Left = &l
	root.Right = &r

	l.Left = &ll
	r.Right = &rr

	rr.Left = &rrl
	rr.Right = &rrr

	fmt.Println("Tree: ", traverseRecursive(&root, ""))
}
