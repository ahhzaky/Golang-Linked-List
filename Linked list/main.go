// https://www.youtube.com/watch?v=8QoynPUY9_8
/*
 Linked list merupakan algoritma untuk mengurutkan array dengan menunjuk pointer sebagai tempat
 penyimpanan array tersebut.
*/

package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// print data array
func (l linkedlist) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println("\n")
}

func (l *linkedlist) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previoustoDelete := l.head
	for previoustoDelete.next.data != value {
		if previoustoDelete.next.next == nil {
			return
		}
		previoustoDelete = previoustoDelete.next

	}
	previoustoDelete.next = previoustoDelete.next.next
	l.length--
}

func main() {

	// input data
	mylist := linkedlist{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 2}
	node4 := &node{data: 11}
	node5 := &node{data: 7}
	node6 := &node{data: 16}

	// add to node
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)

	// mylist.printListData()
	// mylist.deleteWithValue(11)
	// mylist.printListData()

	// fmt.Println(mylist)
}
