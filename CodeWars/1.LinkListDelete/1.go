package main

import "fmt"

type node struct {
  data int
  next *node
}

type list struct{
  head *node
  tail *node
}

func (l *list) Append(val int){
  n := &node {data:val}
  if l.head == nil {
    l.head = n
    l.tail = n
  } else {

  }
}

func main(){
  fmt.Println("hello")
}
