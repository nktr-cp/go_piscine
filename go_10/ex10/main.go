package main

import (
	"fmt"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "2")
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "3")
	piscine.BTreeInsertData(root, "6")
	piscine.BTreeInsertData(root, "5")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "8")
	piscine.BTreeInsertData(root, "9")
	node := piscine.BTreeSearchItem(root, "6")
	fmt.Println("Before delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
	root = piscine.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
}
