/**
    @Author:     ZonzeeLi
    @Project:    lb
    @CreateDate: 10/12/2022
    @UpdateDate: xxx
    @Note:       xxxx
**/

package lb

import "errors"

type RoundRobinBalance struct {
	insts    []*Instance
	curIndex int
}

func init() {
	RegisterBalancer("roundrobin", &RoundRobinBalance{})
}

func (p *RoundRobinBalance) NewBalancer(param ...interface{}) (err error) {
	return
}

func (p *RoundRobinBalance) Add(insts ...*Instance) (err error) {
	if len(insts) <= 0 {
		return errors.New("no instance")
	}
	for _, v := range insts {
		p.insts = append(p.insts, v)
	}
	return
}

func (p *RoundRobinBalance) DoBalance(key ...string) (inst *Instance, err error) {
	insts := p.insts
	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}
	lens := len(insts)
	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens
	return
}
