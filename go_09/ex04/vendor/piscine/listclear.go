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

func ListClear(l *List) {
	for l.Head != nil {
		next := l.Head.Next
		l.Head = nil
		l.Head = next
	}
}
