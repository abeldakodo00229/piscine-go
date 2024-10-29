package piscine

func ListMerge2(l1 *NodeI, l2 *NodeI) {
	p := l2
	for p != nil {
		listPushBack2(l1, p.Data)
		p = p.Next
	}
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	ListMerge2(n1, n2)
	n1 = ListSort(n1)
	return n1
}
