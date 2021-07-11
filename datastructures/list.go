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

//Adds several nodes so that it wont loop several times to get to the end
func (list *StringList) AddNodes(value string) {
	newNode := StringListNode{Next: nil, Value: value}

	//Sets the end node to the new node for the currently last node and then it sets it for the list variable
	(*list.End).Next = &newNode
	list.End = &newNode

}

func (list *StringList) TrimStart(newStartNode *StringListNode) {
	list.Next = newStartNode

	if newStartNode == nil {
		list.Next = nil
		list.End = nil
	}
}
