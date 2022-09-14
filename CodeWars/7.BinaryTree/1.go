package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(k int) {
	if k > n.key {
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.Insert(k)
		}
	} else {
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.Insert(k)
		}
	}
	//node.key == k, is already in tree
}
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if k > n.key {
		//go right
		return n.right.Search(k)

	} else if n.key > k {
		//go left
		return n.left.Search(k)
	}

	return true
}

//Print_Preorder : value - left - right
func (n *Node) Print_Preorder() {
	if n != nil {
		fmt.Println(n.key)
		n.left.Print_Preorder()
		n.right.Print_Preorder()
	}
}

//Print_Inorder : left - value - right
func (n *Node) Print_Inorder() {
	if n != nil {
		n.left.Print_Preorder()
		fmt.Println(n.key)
		n.right.Print_Preorder()
	}
}

//Print_Postorder : left - right - value
func (n *Node) Print_Postorder() {
	if n != nil {
		n.left.Print_Postorder()
		n.right.Print_Postorder()
		fmt.Println(n.key)
	}
}

func main() {
	fmt.Println("hello")
	t := &Node{key: 100}
	t.Insert(203)
	t.Insert(52)
	t.Insert(150)
	t.Insert(310)
	t.Insert(276)
	t.Insert(19)
	t.Insert(76)
	t.Insert(56)
	t.Insert(7)
	t.Insert(24)
	t.Insert(88)
	fmt.Printf("t.key=%d, t.left= %d, t.right=%d \n", t.key, t.left.key, t.right.key)
	// fmt.Println(t.Search(600))
	t.Print_Preorder()
	fmt.Println("--------- PreOrder")
	t.Print_Inorder()
	fmt.Println("--------- InOrder")
	t.Print_Postorder()
	fmt.Println("--------- PostOrder")
}
