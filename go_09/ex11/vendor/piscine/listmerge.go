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

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
	} else if l2.Head != nil {
		l1.Tail.Next = l2.Head
		l1.Tail = l2.Tail
	}
}
