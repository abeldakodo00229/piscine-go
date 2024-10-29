package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	for i := 0; i < pos; i++ {
		if l != nil {
			l = l.Next
		} else {
			return nil
		}
	}
	return l
}
