package main

import (
	"fmt"
)

//Array HashTable
//Bucket

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*Bucket
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	fmt.Printf("%s, %d \n", key, index)
	h.array[index].Insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].Search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].Delete(key)

}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum = sum + int(v)
	}
	fmt.Printf("key:%s,sum=%d, hash=%d \n", key, sum, sum%ArraySize)
	return sum % ArraySize
}

type Node struct {
	key  string
	next *Node
}

type Bucket struct {
	first *Node
	last  *Node
	head  *Node
}

func (b *Bucket) Insert(k string) {
	n := &Node{key: k}
	if b.first == nil {
		b.first = n
		b.last = n
		b.head = n
	} else {
		if b.Search(k) {
			return
		}
		b.last.next = n
		b.last = n
	}
}

func (b *Bucket) Delete(k string) {
	// b.head = b.first
	current := b.first
	var prev *Node = nil
	for current != nil {
		if current.key == k {
			//first node is match
			if b.first == current {
				nextNode := b.first.next
				b.first = nextNode
				b.head = nextNode
				b.last = nextNode
				return
			}
			//if last node
			if current == b.last {
				b.last = prev
				b.last.next = nil
				return
			}
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}
}

func (b *Bucket) Search(k string) bool {
	current := b.first
	for current != nil {
		if current.key == k {
			return true
		}
		current = current.next
	}
	return false
}

func (b *Bucket) Print() {
	current := b.first
	for current != nil {
		fmt.Println(current.key)
		current = current.next
	}
}

func NewHashTable() *HashTable {
	h := &HashTable{}
	for i := range h.array {
		h.array[i] = &Bucket{}
	}
	return h
}

func main() {
	fmt.Println("hello")
	l := Bucket{}
	l.Insert("1")
	l.Insert("2")
	l.Insert("3")
	fmt.Println(l)
	l.Delete("2")
	l.Print()
	fmt.Println(l)
	/* l := Bucket{}
	l.Insert("1")
	l.Insert("2")
	l.Insert("3")
	l.Insert("1")
	l.Print()
	fmt.Println(l.Search("5"))
	fmt.Println(l.Search("4"))
	fmt.Println(l.Search("1"))
	l.Delete("3")
	l.Print()
	fmt.Println("hash func")
	fmt.Println("najam:", hash("najam"))
	fmt.Println("najaf:", hash("najaf"))
	fmt.Println("moh:", hash("moh"))
	fmt.Println("vivi:", hash("vivi")) */
	h := NewHashTable()
	h.Insert("najam")
	fmt.Println(h)
	fmt.Println(h.array[1].head.key)

}
