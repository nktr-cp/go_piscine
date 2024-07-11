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

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Data > elem {
		return BTreeSearchItem(root.Left, elem)
	} else if root.Data < elem {
		return BTreeSearchItem(root.Right, elem)
	} else {
		return root
	}
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}
	f(root.Data)
	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	}
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}
	if node.Parent == nil {
		root = rplc
	}
	if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}
