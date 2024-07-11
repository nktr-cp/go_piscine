package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPopFront(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	toReturn := l.Head.Data
	l.Head = l.Head.Next
	return toReturn
}

func ListPushBack(l *List, data interface{}) {
	toAdd := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = toAdd
		l.Tail = toAdd
	} else {
		l.Tail.Next = toAdd
		l.Tail = toAdd
	}
}

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

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	queue := &List{}
	ListPushBack(queue, root)

	for queue.Head != nil {
		node := ListPopFront(queue).(*TreeNode)
		if _, err := f(node.Data); err != nil {
			return
		}
		if node.Left != nil {
			ListPushBack(queue, node.Left)
		}
		if node.Right != nil {
			ListPushBack(queue, node.Right)
		}
	}
}
