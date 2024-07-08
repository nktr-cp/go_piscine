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

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		return
	}
	prev := l.Head
	for prev.Next != nil {
		if prev.Next.Data == data_ref {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
}
