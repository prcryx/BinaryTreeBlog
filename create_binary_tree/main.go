package main

import "fmt"

type Node struct {
	data  int
	right *Node
	left  *Node
}

func CreateFirstNode(data int) *Node {
	return &Node{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (node *Node) Insert(data int) {
	if node == nil || data == -999 {
		return
	} else {
		if node.data >= data {
			//add data in left node
			if node.left == nil {
				node.left = &Node{
					data:  data,
					left:  nil,
					right: nil,
				}
			} else {
				node.left.Insert(data)
			}
		} else {
			//add data in right node
			if node.right == nil {
				node.right = &Node{
					data:  data,
					left:  nil,
					right: nil,
				}
			} else {
				node.right.Insert(data)
			}
		}
	}
}

func (root *Node) Display() {
	fmt.Printf("    %v    \n", root.data)
	fmt.Printf("   / \\    \n")
	fmt.Printf("  %v   %v   \n", root.left.data, root.right.data)
	fmt.Printf(" / \\        \n")
	fmt.Printf("%v   %v        \n", root.left.left.data, root.left.right.data)
}

func main() {

	root := CreateFirstNode(10)

	root.Insert(20)
	root.Insert(5)

	root.Insert(3)
	root.Insert(-999) // this will return the cursor to previous node
	root.Insert(6)

	root.Display()
}
