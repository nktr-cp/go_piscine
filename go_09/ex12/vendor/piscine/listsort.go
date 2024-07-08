package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func listSize(l *NodeI) int {
	count := 0
	for l != nil {
		count++
		l = l.Next
	}
	return count
}

func ListSort(l *NodeI) *NodeI {
	lstSize := listSize(l)
	for T := 0; T < lstSize; T++ {
		it := l
		for it.Next != nil {
			if it.Data > it.Next.Data {
				it.Data, it.Next.Data = it.Next.Data, it.Data
			}
			it = it.Next
		}
	}
	return l
}
