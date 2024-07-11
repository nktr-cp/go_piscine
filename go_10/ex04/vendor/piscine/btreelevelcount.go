package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else if data > root.Data {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}

	return root
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 1
	left_level := BTreeLevelCount(root.Left) + 1
	right_level := BTreeLevelCount(root.Right) + 1

	if level < left_level {
		level = left_level
	}
	if level < right_level {
		level = right_level
	}

	return level
}
