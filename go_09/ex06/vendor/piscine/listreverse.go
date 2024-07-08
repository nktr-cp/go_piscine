package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
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

func ListReverse(l *List) {
	prev := (*NodeL)(nil)
	cur := l.Head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	cur.Next = prev
	l.Head, l.Tail = l.Tail, l.Head
}
