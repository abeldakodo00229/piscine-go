package piscine

// import "fmt"
/*
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}*/

func ListPushFront(l *List, data interface{}) {
	Node := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = Node
		l.Tail = Node
		return
	}
	Node.Next = l.Head
	l.Head = Node
}

/*
func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}
*/
