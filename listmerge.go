package piscine

func ListMerge(l1 *List, l2 *List) {
	p := l2.Head
	for p != nil {
		ListPushBack(l1, p.Data)
		p = p.Next
	}
	ListClear(l2)
}
