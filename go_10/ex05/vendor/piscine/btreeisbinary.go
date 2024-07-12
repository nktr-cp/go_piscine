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

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Parent != nil {
		if root.Parent.Left == root {
			if root.Data >= root.Parent.Data {
				return false
			}
		} else if root.Parent.Right == root {
			if root.Data <= root.Parent.Data {
				return false
			}
		} else {
			return false
		}
	}

	if root.Left != nil && root.Left.Data >= root.Data {
		return false
	}
	if root.Right != nil && root.Right.Data <= root.Data {
		return false
	}
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
