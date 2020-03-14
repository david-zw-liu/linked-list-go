package linkedlist

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList(nums []int) LinkedList {
	linkedList := LinkedList{}

	head := linkedList.Head
	for _, num := range nums {
		newNode := &Node{Val: num}
		if head == nil {
			linkedList.Head = newNode
		} else {
			head.Next = newNode
		}

		head = newNode
	}

	return linkedList
}

func splitFrontBack(head *Node) (*Node, *Node) {
	if head == nil || head.Next == nil {
		return head, nil
	}

	slow := head
	fast := head.Next
	for (slow.Next != nil) && (fast != nil && fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next
	}
	back := slow.Next
	slow.Next = nil

	return head, back
}

func merge(a, b *Node) *Node {
	// use null head to hold new link
	nullHead := &Node{}
	cursor := nullHead

	// compare and append to cursor node
	for a != nil && b != nil {
		if a.Val < b.Val {
			cursor.Next = a
			a = a.Next
		} else {
			cursor.Next = b
			b = b.Next
		}

		cursor = cursor.Next
		cursor.Next = nil
	}

	// merge remaining nodes
	if a != nil {
		cursor.Next = a
	} else {
		cursor.Next = b
	}

	return nullHead.Next
}

func mergeSort(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	front, back := splitFrontBack(head)
	front = mergeSort(front)
	back = mergeSort(back)

	return merge(front, back)
}

func (linkedList *LinkedList) Sort() {
	linkedList.Head = mergeSort(linkedList.Head)
}
