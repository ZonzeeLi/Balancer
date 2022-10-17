package lb

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	insts []*Instance
}

func init() { //用于BalanceMgr
	RegisterBalancer("random", &RandomBalance{})
}

func (p *RandomBalance) NewBalancer(param ...interface{}) (err error) {
	return
}

func (p *RandomBalance) Add(insts ...*Instance) (err error) {
	if len(insts) <= 0 {
		return errors.New("no instance")
	}
	for _, v := range insts {
		p.insts = append(p.insts, v)
	}
	return
}

func (p *RandomBalance) DoBalance(key ...string) (inst *Instance, err error) { //实现方法
	insts := p.insts
	if len(insts) == 0 {
		err = errors.New("no instances")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens) //用随机算法实现负载均衡
	inst = insts[index]
	return
}
