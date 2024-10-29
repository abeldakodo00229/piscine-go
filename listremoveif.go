package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	p := l.Head
	if p == nil {
		return
	}
	if p.Data == data_ref {
		l.Head = p.Next
		p = nil
	}
	p = l.Head
	precedent := p
	if p != nil {
		for p != nil {
			if p.Data == data_ref {
				tmp := p
				precedent.Next = tmp.Next
				tmp = nil
			}
			if p.Data != data_ref {
				precedent = p
			}
			p = p.Next
		}
	}
}
