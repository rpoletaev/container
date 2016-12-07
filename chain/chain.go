package chain

type Item struct {
	Value interface{}
	Prev  *Item
	Next  *Item
}

type Chain struct {
	first *Item
	last  *Item
}

func (c *Chain) Add(i interface{}) {
	item := &Item{Value: i}

	if c.last == nil {
		c.first = item
		c.last = item
		return
	}

	item.Prev = c.last
	c.last.Next = item
	c.last = item
}

func (c *Chain) Each(fn func(v interface{})) {
	for i := c.first; i != nil; i = i.Next {
		fn(i.Value)
	}
}

func (c *Chain) EachBack(fn func(v interface{})) {
	for i := c.last; i != nil; i = i.Prev {
		fn(i.Value)
	}
}
