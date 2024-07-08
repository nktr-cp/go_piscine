package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l.Data > data_ref {
		n := &NodeI{Data: data_ref}
		n.Next = l
		return n
	}
	cur := l
	for cur.Next != nil {
		if cur.Next.Data > data_ref {
			break
		}
		cur = cur.Next
	}
	n := &NodeI{Data: data_ref}
	n.Next = cur.Next
	cur.Next = n
	return l
}

// 破壊的な操作でよいか？
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	for n2 != nil {
		n1 = SortListInsert(n1, n2.Data)
		n2 = n2.Next
	}
	return n1
}
