/*
@Title : MyWithCancel
@Description :
@Author : 谭靖渝
@Update : 2021/2/28 10:47
*/
package MyContext

import (
	"sync"
)
var judge = 1
var closedchan = make(chan struct{})
type CancelFunc func()

type MyEemCtx struct {
	MyContext
	mu       sync.Mutex
	done     chan struct{}
}

type MyContext interface {
	MyDone()<-chan struct{}
	MyCancel()
}

func (c *MyEemCtx)MyDone()<-chan struct{}  {
	c.mu.Lock()
	if c.done == nil {
		c.done = make(chan struct{})
	}
	if judge==2{
		close(c.done)
	}
	d := c.done
	c.mu.Unlock()
	return d
}

func (c *MyEemCtx)MyCancel()  {
	judge=2
}

func MyWithCancel()(ctx MyEemCtx,cancel CancelFunc)  {
	c:=MyEemCtx{}
	return c, func() { c.MyCancel() }
}


