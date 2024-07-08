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
