/**
    @Author:     ZonzeeLi
    @Project:    balance
    @CreateDate: 10/17/2022
    @UpdateDate: xxx
    @Note:       xxxx
**/

package lb

import (
	"container/heap"
	"errors"
)

type LeastLoadBalance struct {
	pq *h
}

func init() {
	RegisterBalancer("LeastLoad", &LeastLoadBalance{})
}

func (p *LeastLoadBalance) NewBalancer(param ...interface{}) (err error) {
	return
}

func (p *LeastLoadBalance) Add(insts ...*Instance) (err error) {
	if len(insts) <= 0 {
		return errors.New("no instance")
	}
	pq := p.pq
	heap.Init(pq)
	for _, v := range insts {
		heap.Push(pq, v)
	}
	return
}

func (p *LeastLoadBalance) DoBalance(ip ...string) (inst *Instance, err error) {
	pq := p.pq
	if pq.Len() <= 0 {
		err = errors.New("no instance")
		return
	}
	// 如果ip不为空，则创建新的inst也补充进来

	inst = (*pq)[0]
	return inst, nil
}

type h []*Instance

func (h *h) Len() int           { return len(*h) }
func (h *h) Less(i, j int) bool { return (*h)[i].Connections < (*h)[j].Connections }
func (h *h) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *h) Push(v interface{}) {
	*h = append(*h, v.(*Instance))
}

func (h *h) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}
