package main

import "fmt"

type Qeue struct {
	items []*Node
}

func (q *Qeue) Push(k *Node) {
	q.items = append(q.items, k)
}

func (q *Qeue) Pop() *Node {
	k := q.items[0]
	q.items = q.items[1:]
	return k
}

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

func (n *Node) DF_PreOrder() {
	if n != nil {
		fmt.Println(n.key)
		n.left.DF_PreOrder()
		n.right.DF_PreOrder()
	}
}

func (n *Node) DF_InOrder() {
	if n != nil {
		n.left.DF_PreOrder()
		fmt.Println(n.key)
		n.right.DF_PreOrder()
	}

}

func (n *Node) DF_PostOrder() {
	if n != nil {
		n.left.DF_PreOrder()
		n.right.DF_PreOrder()
		fmt.Println(n.key)
	}

}

func (n *Node) BreadthFirst() {
	//level order traversal
	//1.take empty q and intial with root node
	q := Qeue{}
	q.Push(n)
	//2. iterate over q until its not empty
	for len(q.items) != 0 {
		//a.pope item from Qeue
		i := q.Pop()
		//b.prints its value
		fmt.Println(i)
		//c. add its left and right child to Qeue
		if i.left != nil {
			q.Push(i.left)
		}
		if i.right != nil {
			q.Push(i.right)
		}
	}
}

func main() {
	//should work
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
	t.BreadthFirst()
	fmt.Println("$$$$$$$$$ DF ")
	t.DF_PreOrder()
	fmt.Println("--------- Pre ")
	t.DF_InOrder()
	fmt.Println("--------- Inorder ")
	t.DF_PostOrder()
	fmt.Println("--------- Post ")
}
