package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
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

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
		return root
	}
	if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
		return root
	}

	if root.Left == nil {
		temp := root.Right
		if temp != nil {
			temp.Parent = root.Parent
		}
		return temp
	} else if root.Right == nil {
		temp := root.Left
		if temp != nil {
			temp.Parent = root.Parent
		}
		return temp
	}
	successor := root.Right
	for successor.Left != nil {
		successor = successor.Left
	}
	root.Data = successor.Data
	root.Right = BTreeDeleteNode(root.Right, successor)
	return root
}
