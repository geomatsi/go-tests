//
// mix of simple tree examples
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

func subtreesRecursive(root *treeNode, res *[]*treeNode, aux map[string]int8) string {
	var sub string = ""

	if root.Left != nil {
		sub += subtreesRecursive(root.Left, res, aux)
		sub += "L"
	}

	sub += fmt.Sprintf(":%v:", root.Val)

	if root.Right != nil {
		sub += subtreesRecursive(root.Right, res, aux)
		sub += "R"
	}

	if v, ok := aux[sub]; ok {
		if v == 1 {
			*res = append(*res, root)
			aux[sub] = 2
		}
	} else {
		aux[sub] = 1
	}

	return sub
}

func findDuplicateSubtrees(root *treeNode) []*treeNode {
	var aux = make(map[string]int8)
	var res = make([]*treeNode, 0)

	if root != nil {
		subtreesRecursive(root, &res, aux)
	}

	return res

}

func main() {

	// Create tree:
	//           0
	//          / \
	//         1   3
	//        /     \
	//       2       1
	//              / \
	//             2   2
	//

	root := treeNode{0, nil, nil}

	l := treeNode{1, nil, nil}
	r := treeNode{3, nil, nil}

	ll := treeNode{2, nil, nil}
	rr := treeNode{1, nil, nil}

	rrl := treeNode{2, nil, nil}
	rrr := treeNode{2, nil, nil}

	root.Left = &l
	root.Right = &r

	l.Left = &ll
	r.Right = &rr

	rr.Left = &rrl
	rr.Right = &rrr

	fmt.Println("Tree: ", traverseRecursive(&root, ""))

	subTrees := findDuplicateSubtrees(&root)

	fmt.Println("Duplicate subtrees:")
	for _, e := range subTrees {
		fmt.Println("-> ", traverseRecursive(e, ""))
	}

	// Create tree:
	//           0
	//          / \
	//         1   3
	//        /     \
	//       2       1
	//              /
	//             2
	//

	root = treeNode{0, nil, nil}

	l = treeNode{1, nil, nil}
	r = treeNode{3, nil, nil}

	ll = treeNode{2, nil, nil}
	rr = treeNode{1, nil, nil}

	rrl = treeNode{2, nil, nil}

	root.Left = &l
	root.Right = &r

	l.Left = &ll
	r.Right = &rr

	rr.Left = &rrl

	fmt.Println("Tree: ", traverseRecursive(&root, ""))

	subTrees = findDuplicateSubtrees(&root)

	fmt.Println("Duplicate subtrees:")
	for _, e := range subTrees {
		fmt.Println("-> ", traverseRecursive(e, ""))
	}

	// Create tree:
	//           0
	//          / \
	//         1   3
	//        / \   \
	//       2   2   1
	//              / \
	//             2   2
	//

	root = treeNode{0, nil, nil}

	l = treeNode{1, nil, nil}
	r = treeNode{3, nil, nil}

	ll = treeNode{2, nil, nil}
	lr := treeNode{2, nil, nil}
	rr = treeNode{1, nil, nil}

	rrl = treeNode{2, nil, nil}
	rrr = treeNode{2, nil, nil}

	root.Left = &l
	root.Right = &r

	l.Left = &ll
	l.Right = &lr
	r.Right = &rr

	rr.Left = &rrl
	rr.Right = &rrr

	fmt.Println("Tree: ", traverseRecursive(&root, ""))

	subTrees = findDuplicateSubtrees(&root)

	fmt.Println("Duplicate subtrees:")
	for _, e := range subTrees {
		fmt.Println("-> ", traverseRecursive(e, ""))
	}
}
