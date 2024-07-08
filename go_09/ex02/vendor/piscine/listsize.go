package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	toAdd := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = toAdd
		l.Tail = toAdd
	} else {
		toAdd.Next = l.Head
		l.Head = toAdd
	}
}

func ListSize(l *List) int {
	result := 0
	cur := l.Head
	for cur != nil {
		result++
		cur = cur.Next
	}
	return result
}
