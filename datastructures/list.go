package datastructures

//has end node so that it wont need to be looped to
type List struct {
	Next *ListNode
	End  *ListNode
}

type ListNode struct {
	Next  *ListNode
	Value interface{}
}

//Adds several nodes so that it wont loop several times to get to the end
func (list *List) AddNodes(value interface{}) {
	newNode := ListNode{Next: nil, Value: value}

	//Sets the end node to the new node for the currently last node and then it sets it for the list variable
	(*list.End).Next = &newNode
	list.End = &newNode

}

func (list *List) TrimStart(newStartNode *ListNode) {
	list.Next = newStartNode

	if newStartNode == nil {
		list.Next = nil
		list.End = nil
	}
}
