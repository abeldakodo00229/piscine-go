package piscine

func ListReverse(l *List) {
	l2 := &List{}
	a := l.Head
	for a != nil {
		ListPushFront(l2, a.Data)
		a = a.Next
	}
	*l = *l2
}
