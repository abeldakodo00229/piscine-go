package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListCount(l *NodeI) int {
	if l == nil {
		return 0
	}
	count := 0
	head := l
	for head != nil {
		count++
		head = head.Next
	}
	return count
}

func ListSort(l *NodeI) *NodeI {
	p := l
	for i := 0; i < ListCount(l); i++ {
		p = l
		for p.Next != nil {
			if p.Data > p.Next.Data {
				tmp := p.Next.Data
				p.Next.Data = p.Data
				p.Data = tmp
			}
			p = p.Next
		}
	}
	return l
}
