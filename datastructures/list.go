package datastructures

//has end node so that it wont need to be looped to
type StringList struct {
	Next *StringListNode
	End  *StringListNode
}

type StringListNode struct {
	Next  *StringListNode
	Value string
}

func (list *StringList) AddNode(value string) {
	newNode := StringListNode{Next: nil, Value: value}

	//Sets the end node to the new node for the currently last node and then it sets it for the list variable
	if (*list).End != nil {
		(*list).End.Next = &newNode
	} else {
		(*list).Next = &newNode
		(*list).End = &newNode
	}

	list.End = &newNode

}

func (list *StringList) TrimStart(newStartNode *StringListNode) {
	list.Next = newStartNode

	if newStartNode == nil {
		list.Next = nil
		list.End = nil
	}
}

func (list *StringList) RemoveOneNode() {
	//removes the first one and sets list.Next.Next to the new start node
	list.TrimStart(list.Next.Next)
}
