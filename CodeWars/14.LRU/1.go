package main

import (
	"container/list"
	"fmt"
)

type CacheItem struct {
	Key string
	Val string
	// Node *list.Element
}

type Cache struct {
	Cap    int
	Mapper map[string]*list.Element
	Items  *list.List
}

func (c *Cache) Set(key, val string) {
	if el, ok := c.Mapper[key]; ok {
		//already in mapper and list
		c.Items.MoveToFront(el)
	} else {
		//not in map so add it to list
		cItem := CacheItem{Key: key, Val: val}
		el := c.Items.PushFront(cItem)
		c.Mapper[key] = el
		fmt.Println("cache items length:", c.Items.Len())
		l := c.Items.Len()
		if l > c.Cap {
			le := c.Items.Back()
			cc := le.Value.(CacheItem)
			delete(c.Mapper, cc.Key)
			c.Items.Remove(le)
			//XXX: now if we are removing least recently used items from back
			//also needs to delete it from c.Items
			//problem is we are only storing value on linklist item not key.

		}
	}
}

func (c *Cache) Get(key string) string {
	if el, ok := c.Mapper[key]; ok {
		//found element
		c.Items.MoveToFront(el)
		return el.Value.(CacheItem).Val
	}
	return ""
}

func NewCache(k int) *Cache {
	m := make(map[string]*list.Element, k)
	return &Cache{Cap: k, Items: list.New(), Mapper: m}
}

func main() {
	c := NewCache(5)
	c.Set("pk", "Pakistan")

	c.Set("uk", "united kingdom")
	c.Set("uae", "united arab emrates")
	c.Set("usa", "united states of america")
	c.Set("ro", "romania")
	c.Set("ge", "germany")
	fmt.Printf("%#v \n", c)
	fmt.Println(c.Get("uk"))
	fmt.Println("-------Map-----------")
	for k, v := range c.Mapper {
		fmt.Printf("key:%v, value:%v \n", k, v.Value)
	}
	fmt.Println("-------Cache Items----------")
	for e := c.Items.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		fmt.Printf("%#v\n", e.Value)
	}
}
