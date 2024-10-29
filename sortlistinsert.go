package piscine

func listPushBack2(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}
	if l == nil {
		return n
	}

	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	l = listPushBack2(l, data_ref)
	l = ListSort(l)
	return l
}
